package godaddy

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// Config is the godaddy config
type Config struct {
	ttl        int
	weight     int
	priority   int
	port       int
	data       string
	domain     string
	name       string
	recordType string
	apiURL     string
	keyFile    string

	apiKey
	client *resty.Client
}

type apiKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type record struct {
	Data     string `json:"data"`
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Protocol string `json:"protocol"`
	Service  string `json:"service"`
	TTL      int    `json:"ttl"`
	Weight   int    `json:"weight"`
}

type responseRecord struct {
	Name       string `json:"name"`
	Data       string `json:"data"`
	RecordType string `json:"type"`
	TTL        int    `json:"ttl"`
}

func NewGodaddyConfig(cmd *cobra.Command, args []string) *Config {
	var apiKey apiKey
	data, _ := cmd.Flags().GetString("data")
	ttl, _ := cmd.Flags().GetInt("ttl")
	port, _ := cmd.Flags().GetInt("port")
	priority, _ := cmd.Flags().GetInt("priority")
	weight, _ := cmd.Flags().GetInt("weight")
	domain, _ := cmd.Flags().GetString("domain")
	name, _ := cmd.Flags().GetString("name")
	recordType, _ := cmd.Flags().GetString("type")
	apiURL, _ := cmd.Parent().PersistentFlags().GetString("api-url")
	keyFile, _ := cmd.Parent().PersistentFlags().GetString("key-file")
	proxy, _ := cmd.Parent().PersistentFlags().GetString("proxy")

	content, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(content, &apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	client := resty.New()
	if proxy != "" {
		client.SetProxy(proxy)
	}

	return &Config{
		ttl:        ttl,
		port:       port,
		priority:   priority,
		weight:     weight,
		data:       data,
		domain:     domain,
		name:       name,
		recordType: recordType,

		apiURL:  apiURL,
		keyFile: keyFile,

		apiKey: apiKey,
		client: client,
	}
}
