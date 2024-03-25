package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a DNS record",
	Long:  `Use give info to delete a dns record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Del()
	},
}

func init() {
	godaddyCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")
	delCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	delCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	delCmd.Flags().String("name", "", "The DNS record host")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = delCmd.MarkFlagRequired("name")
}
