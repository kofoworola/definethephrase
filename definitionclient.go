package main

import (
	"encoding/json"
	"github.com/kofoworola/definethephrase/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type DefinitionClient struct {
	Phrase, Provider, Definition string
}

func (client *DefinitionClient) CheckDefinition() (bool, string, error) {
	var (
		result string
		err    error
	)
	switch client.Provider {
	case "oxford":
		if strings.Contains(client.Phrase, " ") {
			return false, "You can not get definitons for phrases on Oxford :(", nil
		} else {
			result, err = checkOxford(client.Phrase)
		}
	case "urbandictionary":
		result, err = checkUrbanDictionary(client.Phrase)
	default:
		if strings.Contains(client.Phrase, " ") {
			result, err = checkUrbanDictionary(client.Phrase)
		}else{
			result, err = checkOxford(client.Phrase)
		}
	}
	if err != nil {
		return false, "", err
	}
	if result == ""{
		return false, "No definitions found",nil
	}else{
		return true, result, nil
	}
}

func checkOxford(phrase string) (string, error) {
	req, _ := http.NewRequest("GET", "https://od-api.oxforddictionaries.com/api/v1/entries/en/"+phrase+"/regions=us", nil)
	req.Header.Add("app_id", os.Getenv("OXFORDAPP_ID"))
	req.Header.Add("app_key", os.Getenv("OXFORDAPP_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var response models.OxfordResponse
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(byteArray, &response)
	if err != nil {
		return "", err
	}
	definitions := response.GetDefinitions()
	if len(definitions) < 1 {
		return "", nil
	}
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(definitions)) + " definitions were found for '" + phrase + "'\n")
	for index, item := range definitions {
		sb.WriteString(strconv.Itoa(index+1) + ". " + item + "\n")
	}
	return sb.String(), nil
}

func checkUrbanDictionary(phrase string) (string, error) {
	var response models.UrbanDictionaryResponse
	client := http.Client{}
	resp, err := client.Get("http://urbanscraper.herokuapp.com/define/" + url.QueryEscape(phrase))
	if err != nil {
		return "", err
	}
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(byteArray, &response)
	if err != nil {
		return "", err
	}
	if response.Id == "" {
		return "", nil
	}
	return "The most popular definition of '" + phrase + "' is: \n" + response.Definition, nil
}
