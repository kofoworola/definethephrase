package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const OXFORDENDPOINT = "https://od-api.oxforddictionaries.com/api/v1/entries"

type DefinitionClient struct {
	Phrase, Provider, Definition string
}

func (client *DefinitionClient) CheckDefinition() string {
	switch client.Provider {
	case "oxford":
	default:
		checkOxford(client.Phrase);
	}
	return ""
}

func checkOxford(phrase string) {
	response, err := http.Get(OXFORDENDPOINT + "/en/" + phrase + "/regions=us")
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	} else{
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(data)
	}
}
