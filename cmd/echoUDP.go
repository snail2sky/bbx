/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/snail2sky/bbx/app/echo"

	"github.com/spf13/cobra"
)

// udpCmd represents the udp command
var udpCmd = &cobra.Command{
	Use:   "udp",
	Short: "UDP server",
	Long:  `UDP echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		config := echo.NewConfig(cmd, args)
		config.UDPRun()
	},
}

func init() {
	echoCmd.AddCommand(udpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// udpCmd.PersistentFlags().String("foo", "", "A help for foo")
	udpCmd.Flags().String("host", "0.0.0.0", "ECHO server listen on this address")
	udpCmd.Flags().Uint("port", 6789, "ECHO server listen on this port")
	udpCmd.Flags().Uint("buf-size", 1024, "ECHO server receive buffer size")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// udpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
