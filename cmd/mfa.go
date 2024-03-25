package cmd

import (
	"fmt"
	"github.com/snail2sky/bbx/app/mfa"
	"github.com/spf13/cobra"
	"log"
)

// mfaCmd represents the mfa command
var mfaCmd = &cobra.Command{
	Use:   "mfa",
	Short: "The google authenticator",
	Long:  `Help get google authenticator MFA code`,
	Run: func(cmd *cobra.Command, args []string) {
		secret, _ := cmd.Flags().GetString("secret")
		if secret == "" {
			_, _ = fmt.Scanln(&secret)
		}
		if secret == "" {
			log.Fatalln("Missing secret error")
		}
		otp := mfa.GetTOTPToken(secret)
		fmt.Println(otp)
	},
}

func init() {
	rootCmd.AddCommand(mfaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mfaCmd.PersistentFlags().String("foo", "", "A help for foo")
	mfaCmd.Flags().String("secret", "", "The google secret, If not use this flag, the program will be receive from stdin")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mfaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
