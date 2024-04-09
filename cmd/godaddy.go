package cmd

import (
	"github.com/spf13/cobra"
)

// godaddyCmd represents the godaddy command
var godaddyCmd = &cobra.Command{
	Use:   "godaddy",
	Short: "Manage godaddy domain",
	Long:  `This tool will be add, delete, update, query dns record`,
}

func init() {
	rootCmd.AddCommand(godaddyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	godaddyCmd.PersistentFlags().StringP("key-file", "k", "./godaddy.key",
		`The godaddy api keyfile: {"key": "????", "secret": "????"}`)

	godaddyCmd.PersistentFlags().String("api-url", "https://api.godaddy.com/v1/domains",
		`The godaddy api URL`)

	godaddyCmd.PersistentFlags().String("proxy", "",
		`The proxy url, support only http proxy. such as http://127.0.0.1:1080`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godaddyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
