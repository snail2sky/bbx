package cmd

import (
	"github.com/snail2sky/bbx/app/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Expose the get function as http service",
	Long: `Expose the get function as http service
/get/car/CAR_NAME?num=123851`,
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	serverCmd.Flags().String("host", "0.0.0.0", "The server will be listening on")
	serverCmd.Flags().String("port", "8080", "The server will be listening on")

	serverCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
