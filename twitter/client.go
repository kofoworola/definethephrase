package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
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
	RetweetedObject map[string]interface{} `json:"retweeted_status"`
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
	content := "@" + tweet.User.Handle + " Hi " + name + ". " + response
	tweets := RecursiveSplit(content, 280)
	var (
		returnedTweet *Tweet
		err           error
	)
	for index, item := range tweets {
		if index == 0 {
			returnedTweet, err = SendTweet(item, tweet.IdStr)
			if err != nil {
				fmt.Println("Could not send tweet response: " + err.Error())
			}
		} else {
			returnedTweet, err = SendTweet(item, returnedTweet.IdStr)
			if err != nil {
				fmt.Println("Could not send tweet response: " + err.Error())
			}
		}
	}
}

func SendErrorResponse(tweet *Tweet, response string){
	name := strings.Trim(tweet.User.Name, " ")
	if index := strings.Index(name, " "); index != -1 {
		name = name[:index]
	}
	message := "@"+tweet.User.Handle+" Hi "+ name+ ". "+response
	if _,err := SendTweet(message,tweet.IdStr); err != nil{
		fmt.Println("Error Occurred: "+ err.Error())
	}
}

func SendTweet(tweet string, reply_id string) (*Tweet, error) {
	fmt.Println("Sending tweet as reply to " + reply_id)
	var responseTweet Tweet
	params := url.Values{}
	params.Set("status", tweet)
	params.Set("in_reply_to_status_id", reply_id)
	client := CreateClient()
	resp, err := client.PostForm("https://api.twitter.com/1.1/statuses/update.json", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &responseTweet)
	if err != nil {
		return nil, err
	}
	return &responseTweet, nil
}

func RecursiveSplit(text string, limit int) []string {
	if len(text) <= limit {
		result := []string{text}
		return result
	} else {
		regex := regexp.MustCompile(`(?:[a-zA-Z0-9_!,.])\s(?:[a-zA-Z0-9_!,.])?`)
		occurrences := regex.FindAllStringIndex(text[:limit-3], -1)
		lastOccurrence := occurrences[len(occurrences)-1]
		firstSection := text[:lastOccurrence[0]+1] + "..."
		secondSection := text[lastOccurrence[0]+1:]
		result := []string{firstSection}
		result = append(result, RecursiveSplit(secondSection, limit)...)
		return result
	}
}

