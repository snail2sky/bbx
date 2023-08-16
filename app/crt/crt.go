package crt

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/snail2sky/bbx/types"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CertificateChecker defines the interface for different certificate checking strategies.
type CertificateChecker interface {
	CheckCertificate(certPath, webhookURL string, warningDays int) error
}

// DefaultCertificateChecker implements the default certificate checking strategy.
type DefaultCertificateChecker struct{}

func (d *DefaultCertificateChecker) CheckCertificate(certPath, webhookURL string, warningDays int) error {
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(certData)
	if block == nil || block.Type != "CERTIFICATE" {
		return fmt.Errorf("failed to decode certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	daysUntilExpiration := int(cert.NotAfter.Sub(time.Now()).Hours() / 24)
	if daysUntilExpiration <= warningDays {
		msg := fmt.Sprintf("Certificate at %s will expire in %d days\n", certPath, daysUntilExpiration)

		alert(webhookURL, "证书即将到期", msg)
	}

	return nil
}

// CertificateCheckerFactory creates a CertificateChecker based on the specified strategy.
type CertificateCheckerFactory struct{}

func (f *CertificateCheckerFactory) CreateChecker(strategy string) CertificateChecker {
	switch strategy {
	case "default":
		return &DefaultCertificateChecker{}
	default:
		return &DefaultCertificateChecker{}
	}
}

func Run(data *types.CrtData) {
	certPath := data.CertPath
	warningDays := data.WarningDays
	targetSuffixes := data.TargetSuffixes
	webhookURL := data.WebhookURL
	checker := &DefaultCertificateChecker{}

	fileInfo, err := os.Stat(certPath)
	if err != nil {
		log.Fatal("Error accessing the provided path:", err)
	}

	if fileInfo.IsDir() {
		err := processDirectory(certPath, webhookURL, targetSuffixes, warningDays, checker)
		if err != nil {
			log.Fatal("Error processing directory:", err)
		}
	} else {
		err := checker.CheckCertificate(certPath, webhookURL, warningDays)
		if err != nil {
			log.Fatal("Error processing certificate:", err)
		}
	}
}

func processDirectory(dirPath, webhookURL, targetSuffixes string, warningDays int, checker CertificateChecker) error {
	suffixList := strings.Split(targetSuffixes, ",")

	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !hasSuffix(path, suffixList) {
			return nil
		}

		return checker.CheckCertificate(path, webhookURL, warningDays)
	})
}

func hasSuffix(filename string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(filename, suffix) {
			return true
		}
	}
	return false
}

func alert(url, title, msg string) {
	webhookURL := url

	// 构建要发送的数据
	message := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": fmt.Sprintf("%s: %s", title, msg),
		},
	}

	// 将数据转换为 JSON 格式
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	// 发送 POST 请求
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Alert message sent successfully!", msg)
	} else {
		log.Println("Failed to send alert message. Status code:", resp.StatusCode)
	}
}
