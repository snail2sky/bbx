package cmd

import (
	"github.com/snail2sky/bbx/types"

	"github.com/snail2sky/bbx/app/http_echo"
	"github.com/spf13/cobra"
)

// httpEchoCmd represents the httpEcho command
var httpEchoCmd = &cobra.Command{
	Use:   "httpEcho",
	Short: "HTTP echo server",
	Long:  `A http echo server`,
	Run: func(cmd *cobra.Command, args []string) {
		http_echo.Run(&httpEchoData)
	},
}
var httpEchoData types.HTTPEchoData

func init() {
	rootCmd.AddCommand(httpEchoCmd)
	httpEchoCmd.Flags().StringVar(&httpEchoData.Host, "host", "0.0.0.0", "HTTP echo server listen on this address.")
	httpEchoCmd.Flags().IntVar(&httpEchoData.Port, "port", 8080, "HTTP echo server listen on this port.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpEchoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpEchoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
