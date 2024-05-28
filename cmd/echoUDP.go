package cmd

import (
	"github.com/snail2sky/bbx/app/echo"

	"github.com/spf13/cobra"
)

// echoUDPCmd represents the udp command
var echoUDPCmd = &cobra.Command{
	Use:   "udp",
	Short: "UDP server",
	Long:  `UDP echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		config := echo.NewConfig(cmd, args)
		config.UDPRun()
	},
}

func init() {
	echoCmd.AddCommand(echoUDPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoUDPCmd.PersistentFlags().String("foo", "", "A help for foo")
	echoUDPCmd.Flags().String("host", "0.0.0.0", "ECHO server listen on this address")
	echoUDPCmd.Flags().Uint("port", 6789, "ECHO server listen on this port")
	echoUDPCmd.Flags().Uint("buf-size", 1024, "ECHO server receive buffer size")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoUDPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
