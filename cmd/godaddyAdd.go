package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"
	"github.com/spf13/cobra"
)

// godaddyAddCmd represents the add command
var godaddyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a DNS record",
	Long:  `Use give info to add a dns record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Add()
	},
}

func init() {
	godaddyCmd.AddCommand(godaddyAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// godaddyAddCmd.PersistentFlags().String("foo", "", "A help for foo")
	godaddyAddCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	godaddyAddCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	godaddyAddCmd.Flags().String("name", "", "The DNS record host")
	godaddyAddCmd.Flags().String("data", "", "The DNS record data")
	godaddyAddCmd.Flags().Int("ttl", 600, "The DNS record TLL")
	godaddyAddCmd.Flags().Int("weight", 0, "The DNS record weight")
	godaddyAddCmd.Flags().Int("priority", 0, "The DNS record priority")
	godaddyAddCmd.Flags().Int("port", 1, "The DNS record port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godaddyAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = godaddyAddCmd.MarkFlagRequired("name")
	_ = godaddyAddCmd.MarkFlagRequired("data")
}
