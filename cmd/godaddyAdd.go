package cmd

import (
	"github.com/snail2sky/bbx/app/godaddy"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a DNS record",
	Long:  `Use give info to add a dns record`,
	Run: func(cmd *cobra.Command, args []string) {
		config := godaddy.NewGodaddyConfig(cmd, args)
		config.Add()
	},
}

func init() {
	godaddyCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
	addCmd.Flags().StringP("type", "t", "A", "The DNS record type")
	addCmd.Flags().StringP("domain", "d", "snail2sky.live", "The DNS base domain")
	addCmd.Flags().String("name", "", "The DNS record host")
	addCmd.Flags().String("data", "", "The DNS record data")
	addCmd.Flags().Int("ttl", 600, "The DNS record TLL")
	addCmd.Flags().Int("weight", 0, "The DNS record weight")
	addCmd.Flags().Int("priority", 0, "The DNS record priority")
	addCmd.Flags().Int("port", 1, "The DNS record port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = addCmd.MarkFlagRequired("name")
	_ = addCmd.MarkFlagRequired("data")
}
