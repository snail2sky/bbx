package godaddy

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *Config) Add() {
	url := fmt.Sprintf("%s/%s/records/%s/%s", c.apiURL, c.domain, c.recordType, c.name)
	ssoKey := fmt.Sprintf("sso-key %s:%s", c.apiKey.Key, c.apiKey.Secret)

	request := c.client.R()

	request.SetHeader("accept", "application/json")
	request.SetHeader("Content-Type", "application/json")
	request.SetHeader("Authorization", ssoKey)

	dnsRecord := record{
		Data:     c.data,
		Port:     c.port,
		Priority: c.priority,
		Protocol: "string",
		Service:  "string",
		TTL:      c.ttl,
		Weight:   c.weight,
	}

	body, err := json.Marshal([]record{dnsRecord})
	if err != nil {
		log.Fatalln(err)
	}
	request.SetBody(body)
	response, err := request.Put(url)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(response.Status(), string(response.Body()))
}
