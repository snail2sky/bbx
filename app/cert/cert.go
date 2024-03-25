package cert

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Config 结构体用于存储命令行参数配置
type Config struct {
	certDirList []string
	suffixList  []string
	expireDays  uint
	recursive   bool
}

// NewConfig 方法用于创建 Config 结构体实例
func NewConfig(cmd *cobra.Command, args []string) *Config {
	certDirList, _ := cmd.Flags().GetStringSlice("cert-dir")
	suffixList, _ := cmd.Flags().GetStringSlice("suffix")
	expireDays, _ := cmd.Flags().GetUint("expire")
	recursive, _ := cmd.Flags().GetBool("recursive")
	return &Config{
		certDirList: certDirList,
		suffixList:  suffixList,
		expireDays:  expireDays,
		recursive:   recursive,
	}
}

// walkAndCheck 方法用于遍历证书目录并进行证书校验
func (c *Config) walkAndCheck() {
	var wg sync.WaitGroup

	// 为每个证书目录启动一个 goroutine
	for _, certDir := range c.certDirList {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			c.walkDir(dir)
		}(certDir)
	}

	wg.Wait()
}

// walkDir 方法用于遍历单个证书目录
func (c *Config) walkDir(dir string) {
	dir = expandUser(dir)
	// 定义遍历函数
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		// 检查文件是否为目录
		if info.IsDir() {
			if !c.recursive && path != dir {
				return filepath.SkipDir // 非递归模式下跳过子目录
			}
			return nil
		}

		// 检查文件后缀是否为指定的后缀
		validSuffix := false
		for _, suffix := range c.suffixList {
			if strings.HasSuffix(info.Name(), suffix) {
				validSuffix = true
				break
			}
		}
		if !validSuffix {
			return nil
		}

		// 读取证书文件内容
		certData, err := os.ReadFile(path)
		if err != nil {
			log.Println("Error reading certificate file:", err)
			return nil
		}

		//
		block, rest := pem.Decode(certData)
		for block != nil {
			if block.Type == "CERTIFICATE" {
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					log.Fatalln("Error parsing certificates:", err)
				}
				remainingDays := int(cert.NotAfter.Sub(time.Now()).Hours() / 24)

				//log.Println(cert.Subject.CommonName, remainingDays, c.expireDays)
				//log.Println("---------------------------------------------")

				// 如果剩余有效时间小于指定值，打印警告
				if uint(remainingDays) < c.expireDays {
					log.Printf("Certificate %s in file %s will expire in %d days\n", cert.Subject.CommonName, info.Name(), remainingDays)
				}
			}
			block, rest = pem.Decode(rest)
		}

		return nil
	}

	// 遍历证书目录
	err := filepath.Walk(dir, walkFunc)
	if err != nil {
		log.Println(err)
	}
}

// expandUser 将 ~ 转成真正的家目录的绝对路径写法
func expandUser(path string) string {
	if strings.HasPrefix(path, "~") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = strings.Replace(path, "~", dir, 1)
	}
	return path
}

func (c *Config) Verify() {
	// 创建 Config 结构体实例并解析命令行参数
	// 遍历证书目录并进行证书校验
	c.walkAndCheck()
}
