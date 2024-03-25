/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/snail2sky/bbx/app/echo"

	"github.com/spf13/cobra"
)

// tcpCmd represents the tcp command
var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP server",
	Long:  `TCP echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		config := echo.NewConfig(cmd, args)
		config.TCPRun()
	},
}

func init() {
	echoCmd.AddCommand(tcpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tcpCmd.PersistentFlags().String("foo", "", "A help for foo")
	tcpCmd.Flags().String("host", "0.0.0.0", "ECHO server listen on this address")
	tcpCmd.Flags().Uint("port", 6789, "ECHO server listen on this port")
	tcpCmd.Flags().Uint("buf-size", 1024, "ECHO server receive buffer size")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tcpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
