package alert

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

type FeiShuMsg struct {
	msg        string
	webhookURL string
	atList     []string
}

type FeiShuPayload struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func executeCommand(cmdStr string) (string, error) {
	// 移除方括号
	cmdStr = strings.Trim(cmdStr, "[]")
	cmdParts := strings.Fields(cmdStr)
	if len(cmdParts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func NewFeiShuMsg(cmd *cobra.Command, args []string) *FeiShuMsg {
	var msg string
	if len(args) > 0 {
		// 检查是否包含方括号
		if strings.HasPrefix(args[0], "[") && strings.HasSuffix(args[0], "]") {
			output, err := executeCommand(args[0])
			if err != nil {
				log.Printf("Command execution failed: %v\n", err)
				msg = fmt.Sprintf("Command execution failed: %v", err)
			} else {
				msg = output
			}
		} else {
			msg = strings.Join(args, "\n")
		}
	}

	webhookURL, _ := cmd.Flags().GetString("webhook-url")
	atList, _ := cmd.Flags().GetStringSlice("at")
	return &FeiShuMsg{
		msg:        msg,
		webhookURL: webhookURL,
		atList:     atList,
	}
}

func (f *FeiShuMsg) Send() {
	var atListMsg = make([]string, len(f.atList))
	request := resty.New().R()

	for i, at := range f.atList {
		atListMsg[i] = fmt.Sprintf(`<at user_id="%s"></at>`, at)
	}
	atMsg := strings.Join(atListMsg, "")

	payload := FeiShuPayload{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: atMsg + f.msg,
		},
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Failed to marshal JSON:", err)
	}

	request.SetHeader("Content-Type", "application/json")
	request.SetBody(jsonBody)

	response, err := request.Post(f.webhookURL)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}

	log.Println(response.Status(), string(response.Body()))
}
