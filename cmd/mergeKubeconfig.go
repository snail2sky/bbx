package cmd

import (
	"github.com/snail2sky/bbx/app/merge"
	"github.com/spf13/cobra"
)

// mergeKubeConfigCmd represents the kubeconfig command
var mergeKubeConfigCmd = &cobra.Command{
	Use:   "kubeconfig <merged-file>",
	Short: "merge kubeconfig",
	Long:  `merge kubeconfig to single file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := merge.NewKubeConfig(cmd, args)
		config.Merge()
	},
}

func init() {
	mergeCmd.AddCommand(mergeKubeConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeKubeConfigCmd.PersistentFlags().String("foo", "", "A help for foo")
	mergeKubeConfigCmd.Flags().StringP("config-dir", "c", "./config-dir", "The directory where all kubeconfig files are located")
	mergeKubeConfigCmd.Flags().String("suffix", ".yaml", "The suffix name of the kubeconfig file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mergeKubeConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
