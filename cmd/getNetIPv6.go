package cmd

import (
	"github.com/snail2sky/bbx/app/get/net"

	"github.com/spf13/cobra"
)

// getNetIPv6Cmd represents the ipv6 command
var getNetIPv6Cmd = &cobra.Command{
	Use:   "ipv6",
	Short: "Get IPv6 network addresses",
	Long:  `Get IPv6 address, local or external IP addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		config := net.NewIPv6Config(cmd)
		config.GetIPv6()
	},
}

func init() {
	getNetCmd.AddCommand(getNetIPv6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getNetIPv6Cmd.PersistentFlags().String("foo", "", "A help for foo")
	getNetIPv6Cmd.Flags().String("where", "", "where to find the IP address: '', 'local' or 'all'")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getNetIPv6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
