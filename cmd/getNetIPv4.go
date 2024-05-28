package cmd

import (
	"github.com/snail2sky/bbx/app/get/net"

	"github.com/spf13/cobra"
)

// getNetIPv4Cmd represents the ipv4 command
var getNetIPv4Cmd = &cobra.Command{
	Use:   "ipv4",
	Short: "Get IPv4 address",
	Long:  `Get IPv4 address, local or external IP addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		config := net.NewIPv4Config(cmd)
		config.GetIPv4()
	},
}

func init() {
	getNetCmd.AddCommand(getNetIPv4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNetIPv4Cmd.PersistentFlags().String("foo", "", "A help for foo")
	getNetIPv4Cmd.Flags().String("where", "", "where to find the IP address: '', 'local' or 'all'")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNetIPv4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
