package cmd

import (
	"github.com/snail2sky/bbx/app/echo"

	"github.com/spf13/cobra"
)

// echoTCPCmd represents the tcp command
var echoTCPCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP server",
	Long:  `TCP echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		config := echo.NewConfig(cmd, args)
		config.TCPRun()
	},
}

func init() {
	echoCmd.AddCommand(echoTCPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoTCPCmd.PersistentFlags().String("foo", "", "A help for foo")
	echoTCPCmd.Flags().String("host", "0.0.0.0", "ECHO server listen on this address")
	echoTCPCmd.Flags().Uint("port", 6789, "ECHO server listen on this port")
	echoTCPCmd.Flags().Uint("buf-size", 1024, "ECHO server receive buffer size")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoTCPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
