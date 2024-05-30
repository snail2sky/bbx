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
	Run: func(cmd *cobra.Command, args []string) {
		num := cmd.Flag("num").Value.String()
		password, err := car.GetArrizoPassword(num)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(password)
		}
	},
}

func init() {
	getCarCmd.AddCommand(getCarAI8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCarAI8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	getCarAI8Cmd.Flags().StringP("num", "n", "", "Serial number of project")
	_ = getCarAI8Cmd.MarkFlagRequired("num")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCarAI8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
