package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/kofoworola/definethephrase/redisdb"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var endpoint = "https://api.twitter.com/1.1/account_activity"

//Check if all credentials needed were successfully passed
func HasEnvVariables() bool {
	return os.Getenv("CONSUMER_KEY") != "" && os.Getenv("CONSUMER_SECRET") != "" &&
		os.Getenv("ACCESS_TOKEN") != "" && os.Getenv("ACCESS_SECRET") != "" &&
		os.Getenv("WEBHOOK_ENV") != ""
}

//Get webhooks
func GetWebhook() {
	conn := redisdb.GetPool().Get()
	defer conn.Close()
	if !HasEnvVariables() {
		panic("Missing required environment variable")
	}
	fmt.Println("Getting webhooks...")
	client := CreateClient()
	path := endpoint + "/all/" + os.Getenv("WEBHOOK_ENV") + "/webhooks.json"
	resp, _ := client.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	hasId := false
	for _, item := range data {
		if item["valid"].(bool) {
			conn.Do("SET", "webhook_id", item["id"].(string))
			fmt.Println("Webhook id of " + item["id"].(string) + " exists.")
			hasId = true
			break
		}else {
			DeleteWebhook(item["id"].(string))
		}
	}
	if !hasId {
		RegisterWebhook()
	}
}

func RegisterWebhook() {
	fmt.Println("Registering webhook...")
	if !HasEnvVariables() {
		panic("Missing required environment variable")
	}
	if os.Getenv("APP_URL") == "" {
		panic("missing app url")
	}

	httpClient := CreateClient()

	path := endpoint + "/all/" + os.Getenv("WEBHOOK_ENV") + "/webhooks.json"
	values := url.Values{}
	values.Set("url", "https://"+os.Getenv("APP_URL")+"/twitter/webhook")
	resp, _ := httpClient.PostForm(path, values)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	conn := redisdb.GetPool().Get()
	conn.Do("SET", "webhook_id", data["id"].(string))
	fmt.Println("Webhook id of " + data["id"].(string) + " has been registered")
}

func SubscribeWebhook() {
	fmt.Println("Subscribing webapp...")
	if !HasEnvVariables() {
		panic("Missing Environment Variables")
	}
	client := CreateClient()
	path := endpoint + "/all/" + os.Getenv("WEBHOOK_ENV") + "/subscriptions.json"
	resp, _ := client.PostForm(path, nil)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("Subscribed successfully")
	} else {
		fmt.Println("Could not subscribe the webhook. Response below:")
		fmt.Println(string(body))
	}
}

func DeleteWebhook(id string) {
	fmt.Println("Deleting webhook...")
	if !HasEnvVariables() {
		panic("Missing required environment variable")
	}
	client := CreateClient()
	path := endpoint + "/all/" + os.Getenv("WEBHOOK_ENV") + "/webhooks/" + id + ".json"
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("Deleted webhook " + id)
	} else {
		fmt.Println("An Error ocuured. Response below: ")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
