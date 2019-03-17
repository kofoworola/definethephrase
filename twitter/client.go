package twitter

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/url"
	"os"
)

func RegisterWebhook() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")
	//fmt.Println(consumerSecret)
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}
	if os.Getenv("WEBHOOK_ENV") == "" || os.Getenv("APP_URL") == ""{
		panic("missing app url or web env")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/1.1/account_activity/all/"+os.Getenv("WEBHOOK_ENV")+"/webhooks.json"
	values := url.Values{}
	values.Set("url", "https://"+os.Getenv("APP_URL")+"/twitter/webhook")
	resp, _ := httpClient.PostForm(path, values)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))
}
