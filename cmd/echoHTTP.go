package cmd

import (
	"github.com/snail2sky/bbx/app/echo"

	"github.com/spf13/cobra"
)

// echoHTTPCmd represents the http command
var echoHTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP server",
	Long:  `HTTP echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		config := echo.NewConfig(cmd, args)
		config.HTTPRun()
	},
}

func init() {
	echoCmd.AddCommand(echoHTTPCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoHTTPCmd.PersistentFlags().String("foo", "", "A help for foo")
	echoHTTPCmd.Flags().String("host", "0.0.0.0", "ECHO server listen on this address")
	echoHTTPCmd.Flags().Uint("port", 6789, "ECHO server listen on this port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoHTTPCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
