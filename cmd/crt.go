package cmd

import (
	"github.com/snail2sky/bbx/app/crt"
	"github.com/snail2sky/bbx/types"
	"github.com/spf13/cobra"
)

// crtCmd represents the crt command
var crtCmd = &cobra.Command{
	Use:   "crt",
	Short: "Certificate Verification Tool",
	Long:  `Detect the certificate file in the specified directory, and send a message reminder to Feishu.`,
	Run: func(cmd *cobra.Command, args []string) {
		crt.Run(&crtData)
	},
}

var crtData types.CrtData

func init() {
	rootCmd.AddCommand(crtCmd)
	crtCmd.Flags().StringVar(&crtData.CertPath, "cert-path", "./crts", "Path to certificate file or directory")
	crtCmd.Flags().IntVar(&crtData.WarningDays, "warning-days", 30, "Number of days before expiration to show warning")
	crtCmd.Flags().StringVar(&crtData.TargetSuffixes, "target-suffixes", ".crt,.pem", "Comma-separated list of target certificate file suffixes")
	crtCmd.Flags().StringVar(&crtData.WebhookURL, "webhook-url", "https://open.feishu.cn/open-apis/bot/v2/hook/d24ce276-47d4-4200-9a4a-38e3669e8d4a", "FeiShu webhook url")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
