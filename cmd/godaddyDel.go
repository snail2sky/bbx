package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// godaddyDelCmd represents the del command
var godaddyDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a DNS record",
	Long:  `Use give info to delete a dns record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Del()
	},
}

func init() {
	godaddyCmd.AddCommand(godaddyDelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// godaddyDelCmd.PersistentFlags().String("foo", "", "A help for foo")
	godaddyDelCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	godaddyDelCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	godaddyDelCmd.Flags().String("name", "", "The DNS record host")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godaddyDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = godaddyDelCmd.MarkFlagRequired("name")
}
