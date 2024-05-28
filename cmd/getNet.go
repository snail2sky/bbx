package cmd

import (
	"github.com/spf13/cobra"
)

// getNetCmd represents the net command
var getNetCmd = &cobra.Command{
	Use:   "net",
	Short: "Get net info",
	Long:  `Get net info, such as ipv4 address`,
}

func init() {
	getCmd.AddCommand(getNetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
