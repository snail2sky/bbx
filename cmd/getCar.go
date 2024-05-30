package cmd

import (
	"github.com/spf13/cobra"
)

// getCarCmd represents the car command
var getCarCmd = &cobra.Command{
	Use:   "car",
	Short: "Get car Project password",
	Long:  `Get car Project password`,
}

func init() {
	getCmd.AddCommand(getCarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
