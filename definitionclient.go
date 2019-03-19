package main

import (
	"encoding/json"
	"fmt"
	"github.com/kofoworola/definethephrase/models"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type DefinitionClient struct {
	Phrase, Provider, Definition string
}

func (client *DefinitionClient) CheckDefinition() string {
	var result string
	switch client.Provider {
	case "oxford":
		result,_ = checkOxford(client.Phrase)
	}
	return result
}

func checkOxford(phrase string) (string, error) {
	req, _ := http.NewRequest("GET", "https://od-api.oxforddictionaries.com/api/v1/entries/en/"+phrase+"/regions=us", nil)
	req.Header.Add("app_id", os.Getenv("OXFORDAPP_ID"))
	req.Header.Add("app_key", os.Getenv("OXFORDAPP_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return "",err
	}
	var response models.OxfordResponse
	byteArray, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteArray, &response)
	if err != nil{
		return "",err
	}
	definitions := response.GetDefinitions()
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(definitions)) + " definitions were found for '"+ phrase + "'\n")
	for index,item := range definitions{
		sb.WriteString(strconv.Itoa(index+1)+ ". " + item +"\n")
	}
	return sb.String(),error(nil)
}
