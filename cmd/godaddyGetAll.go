package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// getAllCmd represents the getAll command
var getAllCmd = &cobra.Command{
	Use:   "get-all",
	Short: "Get all DNS record info",
	Long:  `Get all DNS record info`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.GetAll()
	},
}

func init() {
	godaddyCmd.AddCommand(getAllCmd)

	// Here you will define your flags and configuration settings.
	getAllCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
