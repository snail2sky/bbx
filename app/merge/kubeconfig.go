package merge

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type KubeConfig struct {
	configDir  string
	mergedFile string
	suffix     string
}

//var kubeConfigDir = flag.String("kubeConfigDir", "./config", "The kube config dir will be merged.")
//var suffix = flag.String("suffix", ".yaml", "Kube config file suffix.")
//var mergeFile = flag.String("mergeFile", "merged.yaml", "Kube config merged file.")

func NewKubeConfig(cmd *cobra.Command, args []string) *KubeConfig {
	mergeFile := args[0]
	configDir, _ := cmd.Flags().GetString("config-dir")
	suffix, _ := cmd.Flags().GetString("suffix")
	return &KubeConfig{
		configDir:  configDir,
		suffix:     suffix,
		mergedFile: mergeFile,
	}
}

func (k *KubeConfig) Merge() {
	combinedConfig := api.NewConfig()

	err := filepath.Walk(k.configDir, func(path string, info os.FileInfo, err error) error {
		singleConfig := api.NewConfig()
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), k.suffix) {
			config, err := clientcmd.LoadFromFile(path)
			if err != nil {
				return err
			}

			clusterName := strings.TrimSuffix(info.Name(), k.suffix)
			userName := clusterName + "-admin"
			// contextName := clusterName + "-context"

			for clusterKey, cluster := range config.Clusters {
				delete(singleConfig.Clusters, clusterKey)
				singleConfig.Clusters[clusterName] = cluster
			}

			for authKey, authInfo := range config.AuthInfos {
				delete(singleConfig.AuthInfos, authKey)
				singleConfig.AuthInfos[userName] = authInfo
			}

			for contextKey, context := range config.Contexts {
				delete(singleConfig.Contexts, contextKey)
				for clusterName := range singleConfig.Clusters {
					context.Cluster = clusterName
				}
				for authKey := range singleConfig.AuthInfos {
					context.AuthInfo = authKey
				}
				singleConfig.Contexts[fmt.Sprintf("%s@%s", userName, clusterName)] = context
			}
			// merge
			for clusterKey, cluster := range singleConfig.Clusters {
				combinedConfig.Clusters[clusterKey] = cluster
			}

			for authKey, authInfo := range singleConfig.AuthInfos {
				combinedConfig.AuthInfos[authKey] = authInfo
			}

			for contextKey, context := range singleConfig.Contexts {
				log.Println(contextKey)
				combinedConfig.Contexts[contextKey] = context
			}
		}

		return nil
	})

	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	combinedConfigBytes, err := clientcmd.Write(*combinedConfig)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile(k.mergedFile, combinedConfigBytes, 0644)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	log.Printf("Combined kubeconfig file created: %s", k.mergedFile)
}
