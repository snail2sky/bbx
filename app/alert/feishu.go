package alert

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type FeiShuMsg struct {
	msg        string
	webhookURL string
	atList     []string
}

func NewFeiShuMsg(cmd *cobra.Command, args []string) *FeiShuMsg {
	msg := args[0]
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
	msg := f.msg

	for i, at := range f.atList {
		atListMsg[i] = fmt.Sprintf(`<at user_id=\"%s\"></at>`, at)
	}
	atMsg := strings.Join(atListMsg, "")

	body := fmt.Sprintf(`{
				"msg_type": "text",
				"content": {
					"text": "%s %s"
				}
			}`, atMsg, msg)

	request.SetHeader("Content-Type", "application/json")
	request.SetBody(body)

	response, err := request.Post(f.webhookURL)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response.Status(), string(response.Body()))
}
