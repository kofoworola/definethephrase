package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type WebhookLoad struct {
	UserId           string  `json:"for_user_id"`
	TweetCreateEvent []Tweet `json:"tweet_create_events"`
}
type Tweet struct {
	Id    int64
	IdStr string `json:"id_str"`
	User  User
	Text  string
}
type User struct {
	Id     int64
	IdStr  string `json:"id_str"`
	Name   string
	Handle string `json:"screen_name"`
}

//Create http client for twitter requests
func CreateClient() *http.Client {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))

	return config.Client(oauth1.NoContext, token)
}

func SendResponse(tweet *Tweet, response string) {
	fmt.Println("Sending Response")
	time.Sleep(2 * time.Second)
	name := strings.Trim(tweet.User.Name, " ")
	if index := strings.Index(name, " "); index != -1 {
		name = name[:index]
	}
	pre := "@" + tweet.User.Handle + " Hi " + name + ". "
	var content string
	if len(response) < 280-len(pre) {
		content = response
	}
	var tweets []string
	tweets = append(tweets, content)
	for index, content := range tweets {
		var (
			returnedTweet *Tweet
			err           error
		)
		if index == 0 {
			returnedTweet, err = SendTweet(pre + content, tweet.IdStr)
			if err != nil {
				panic("Could not send tweet response: " + err.Error())
			}
		} else {
			returnedTweet, err = SendTweet(content, returnedTweet.IdStr)
			if err != nil {
				panic("Could not send tweet response: " + err.Error())
			}
		}
	}
}

func SendTweet(tweet string, reply_id string) (*Tweet, error) {
	fmt.Println("Sending tweet as reply to " + reply_id)
	var responseTweet Tweet
	params := url.Values{}
	params.Set("status",tweet)
	params.Set("in_reply_to_status_id",reply_id)
	client := CreateClient()
	resp, err := client.PostForm("https://api.twitter.com/1.1/statuses/update.json",params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &responseTweet)
	if err != nil{
		return  nil,err
	}
	return &responseTweet, nil
}

func BreakTest(text string, limit int) []string{
	result := make([]string,1)
	if len(text) < limit{
		result = append(result,text)
	}
	return result
}
