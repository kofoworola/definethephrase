package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kofoworola/definethephrase/twitter"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type crcResponse struct {
	ResponseToken string `json:"response_token"`
}

func CrcCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("Crc check occurred at %s\n", time.Now().String())
	token := request.URL.Query()["crc_token"]
	if len(token) < 1 {
		fmt.Println("No crc_token given")
		fmt.Fprintf(writer, "No token given")
		return
	}
	h := hmac.New(sha256.New, []byte(os.Getenv("CONSUMER_SECRET")))
	h.Write([]byte(token[0]))
	encoded := base64.StdEncoding.EncodeToString(h.Sum(nil))
	response, _ := json.Marshal(crcResponse{ResponseToken: "sha256=" + encoded})
	fmt.Fprintf(writer, string(response))
}

func WebhookHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler called")
	body, _ := ioutil.ReadAll(request.Body)
	var load twitter.WebhookLoad
	err := json.Unmarshal(body, &load)
	if err != nil {
		fmt.Println("An error occured: " + err.Error())
	}
	if len(load.TweetCreateEvent) < 1 || load.UserId == load.TweetCreateEvent[0].User.IdStr {
		return
	}
	go parseTweet(load.TweetCreateEvent[0])
}

func parseTweet(tweet twitter.Tweet) {
	reg := regexp.MustCompile("@" + os.Getenv("HANDLE") + " (the|this) (word|phrase)")
	indices := reg.FindStringIndex(tweet.Text)
	if len(indices) != 2 {
		fmt.Println("Could not understand")
		return
	}
	phrase := strings.Trim(tweet.Text[indices[1]:], " ")
	client := DefinitionClient{Phrase: phrase, Provider: "oxford"}
	//Reply after 2 seconds to prevent spam marks
	time.Sleep(2 * time.Second)
	twitter.SendResponse(&tweet,client.CheckDefinition())
}
