package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/kofoworola/definethephrase/redisdb"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var endpoint = "https://api.twitter.com/1.1/account_activity"

//Check if all credentials needed were successfully passed
func hasEnvVariables() bool {
	return os.Getenv("CONSUMER_KEY") != "" && os.Getenv("CONSUMER_SECRET") != "" &&
		os.Getenv("ACCESS_TOKEN") != "" && os.Getenv("ACCESS_SECRET") != "" &&
		os.Getenv("WEBHOOK_ENV") != ""
}

//Create http client for twitter requests
func createClient() *http.Client {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))

	return config.Client(oauth1.NoContext, token)
}

//Get webhooks
func GetWebhook() {
	conn := redisdb.GetPool().Get()
	defer conn.Close()
	if !hasEnvVariables() {
		panic("Missing required environment variable")
	}
	fmt.Println("Getting webhooks...")
	client := createClient()
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
			fmt.Println("Webhook id of" + item["id"].(string) + " exists.")
			hasId = true
			break
		}
	}
	if !hasId {
		RegisterWebhook()
	}
}

func RegisterWebhook() {
	fmt.Println("Registering webhook...")
	if !hasEnvVariables() {
		panic("Missing required environment variable")
	}
	if os.Getenv("APP_URL") == "" {
		panic("missing app url")
	}

	httpClient := createClient()

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
	fmt.Println("Webhook id of" + data["id"].(string) + "has been registered")
}

func SubscribeWebhook() {
	fmt.Println("Subscribing webapp...")
	if !hasEnvVariables() {
		panic("Missing Environment Variables")
	}
	client := createClient()
	path := endpoint + "/all/" + os.Getenv("WEBHOOK_ENV") + "/subscriptions.json"
	resp, _ := client.PostForm(path, nil)
	body,_ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("Subscribed successfully")
	} else{
		fmt.Println("Could not subscribe the webhook. Response below:")
		fmt.Println(string(body))
	}
}
