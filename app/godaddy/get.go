package godaddy

import (
	"fmt"
	"log"
)

func (c *Config) Get() {
	url := fmt.Sprintf("%s/%s/records/%s/%s", c.apiURL, c.domain, c.recordType, c.name)
	ssoKey := fmt.Sprintf("sso-key %s:%s", c.apiKey.Key, c.apiKey.Secret)
	request := c.client.R()

	request.SetHeader("accept", "application/json")
	request.SetHeader("Content-Type", "application/json")
	request.SetHeader("Authorization", ssoKey)

	response, err := request.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(response.Status(), string(response.Body()))
}
