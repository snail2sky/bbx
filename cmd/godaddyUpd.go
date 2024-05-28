package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// godaddyUpdCmd represents the upd command
var godaddyUpdCmd = &cobra.Command{
	Use:   "upd",
	Short: "Update a DNS record",
	Long:  `Use command args update DNS record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Upd()
	},
}

func init() {
	godaddyCmd.AddCommand(godaddyUpdCmd)

	godaddyUpdCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	godaddyUpdCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	godaddyUpdCmd.Flags().String("name", "", "The DNS record host")
	godaddyUpdCmd.Flags().String("data", "", "The DNS record data")
	godaddyUpdCmd.Flags().Int("ttl", 600, "The DNS record TLL")
	godaddyUpdCmd.Flags().Int("weight", 0, "The DNS record weight")
	godaddyUpdCmd.Flags().Int("priority", 0, "The DNS record priority")
	godaddyUpdCmd.Flags().Int("port", 1, "The DNS record port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godaddyAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = godaddyUpdCmd.MarkFlagRequired("host")
	_ = godaddyUpdCmd.MarkFlagRequired("data")
}
