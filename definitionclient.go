package main

import (
	"encoding/json"
	"fmt"
	"github.com/kofoworola/definethephrase/models"
	"io/ioutil"
	"net/http"
)

const OXFORDENDPOINT = "https://od-api.oxforddictionaries.com/api/v1/entries"
const OXFORDAPPID = "86de2890"
const OXFORDAPPKEY = "1f8f224fce906820fee2f4b19bf2b0d2"

type DefinitionClient struct {
	Phrase, Provider, Definition string
}

func (client *DefinitionClient) CheckDefinition() string {
	var result string
	switch client.Provider {
	case "oxford":
		result,_ = checkOxford(client.Phrase);
	}
	return result
}

func checkOxford(phrase string) (string, error) {
	req, _ := http.NewRequest("GET", OXFORDENDPOINT+"/en/"+phrase+"/regions=us", nil)
	req.Header.Add("app_id", OXFORDAPPID)
	req.Header.Add("app_key", OXFORDAPPKEY)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return "",err
	}
	var response models.OxfordResponse
	byteArray, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(byteArray, &response)
	fmt.Println(response.GetDefinitions())
	return "",error(nil)
}
