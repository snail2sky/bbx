package cmd

import (
	"fmt"
	"github.com/snail2sky/bbx/app/get/car"
	"github.com/spf13/cobra"
)

// getCarAI8Cmd represents the ai8 command
var getCarAI8Cmd = &cobra.Command{
	Use:   "ai8",
	Short: "Get Arrizo 8 project password",
	Long:  `Get Arrizo 8 project password`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		password := car.GetArrizoPassword(args[0])
		fmt.Println(password)
	},
}

func init() {
	getCarCmd.AddCommand(getCarAI8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCarAI8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCarAI8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
