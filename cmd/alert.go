package cmd

import (
	"github.com/spf13/cobra"
)

// alertCmd represents the alert command
var alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "A alert tool",
	Long:  `Send alert message use some tool, such as feishu`,
}

func init() {
	rootCmd.AddCommand(alertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
