package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"

	"github.com/spf13/cobra"
)

// godaddyGetCmd represents the get command
var godaddyGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get DNS record info",
	Long:  `Get DNS record info`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Get()
	},
}

func init() {
	godaddyCmd.AddCommand(godaddyGetCmd)

	godaddyGetCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	godaddyGetCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	godaddyGetCmd.Flags().String("name", "", "The DNS record host")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// godaddyDelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = godaddyGetCmd.MarkFlagRequired("name")
}
