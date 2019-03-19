package twitter

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCanTweet(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}
	tweet, err := SendTweet("Hello World!", "")
	if err != nil {
		t.Error("An error occured: " + err.Error())
	}
	if tweet.Id == 0 {
		t.Error("Tweet not created")
	}
	fmt.Println(tweet)
}
