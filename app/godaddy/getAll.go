package godaddy

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *Config) GetAll() {
	var recordList []responseRecord
	url := fmt.Sprintf("%s/%s/records", c.apiURL, c.domain)
	ssoKey := fmt.Sprintf("sso-key %s:%s", c.apiKey.Key, c.apiKey.Secret)
	request := c.client.R()

	request.SetHeader("accept", "application/json")
	request.SetHeader("Content-Type", "application/json")
	request.SetHeader("Authorization", ssoKey)

	response, err := request.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(response.Body(), &recordList)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(response.Status())
	for _, record := range recordList {
		fmt.Printf("%#v\n", record)
	}
}
