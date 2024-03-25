package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// updCmd represents the upd command
var updCmd = &cobra.Command{
	Use:   "upd",
	Short: "Update a DNS record",
	Long:  `Use command args update DNS record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Upd()
	},
}

func init() {
	godaddyCmd.AddCommand(updCmd)

	updCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	updCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	updCmd.Flags().String("name", "", "The DNS record host")
	updCmd.Flags().String("data", "", "The DNS record data")
	updCmd.Flags().Int("ttl", 600, "The DNS record TLL")
	updCmd.Flags().Int("weight", 0, "The DNS record weight")
	updCmd.Flags().Int("priority", 0, "The DNS record priority")
	updCmd.Flags().Int("port", 1, "The DNS record port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = updCmd.MarkFlagRequired("host")
	_ = updCmd.MarkFlagRequired("data")
}
