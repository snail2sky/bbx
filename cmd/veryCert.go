package cmd

import (
	"github.com/snail2sky/bbx/app/cert"

	"github.com/spf13/cobra"
)

// verifyCertCmd represents the cert command
var verifyCertCmd = &cobra.Command{
	Use:   "cert",
	Short: "Verify cert valid",
	Long:  `Verify cert valid and give Expiration`,
	Run: func(cmd *cobra.Command, args []string) {
		config := cert.NewConfig(cmd, args)
		config.Verify()
	},
}

func init() {
	verifyCmd.AddCommand(verifyCertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCertCmd.PersistentFlags().String("foo", "", "A help for foo")
	verifyCertCmd.Flags().StringSlice("cert-dir", []string{"./cert-dir"}, "The cert directory will be verify")
	verifyCertCmd.Flags().StringSlice("suffix", []string{"cert", "crt", "pem"}, "The cert suffix list")
	verifyCertCmd.Flags().Uint("expire", 30, "The cert expire days")
	verifyCertCmd.Flags().Bool("recursive", true, "Whether to recursively traverse the directory specified by cert-dir")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
