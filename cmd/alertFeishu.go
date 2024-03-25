package cmd

import (
	"github.com/snail2sky/bbx/app/alert"

	"github.com/spf13/cobra"
)

// feishuCmd represents the feishu command
var feishuCmd = &cobra.Command{
	Use:   "feishu '<message>'",
	Short: "Send message to group",
	Long:  `Send message to group, support @ somebody`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg := alert.NewFeiShuMsg(cmd, args)
		msg.Send()
	},
}

func init() {
	alertCmd.AddCommand(feishuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feishuCmd.PersistentFlags().String("foo", "", "A help for foo")
	feishuCmd.Flags().StringP("webhook-url", "u", "", "The feishu webhook url")
	feishuCmd.Flags().StringSlice("at", nil, `The message will be at someone, value is user_id list
support at many people, such as: ou_18ea???,ou_1823???`)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feishuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	_ = feishuCmd.MarkFlagRequired("webhook-url")
}
