package twitter

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"regexp"
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

func TestCanFindBreakPoint(t *testing.T){
	sentence := "I am Alive "
	regex := regexp.MustCompile(`(?:[a-zA-Z0-9_!,.])\s(?:[a-zA-Z0-9_!,.])?`)
	value := regex.FindAllStringIndex(sentence,-1)
	//Get last match position of reg expression which should be index of e to ' '+1
	//Check if it matches
	last := value[len(value)-1]
	if letter := string(sentence[last[1]-1]); letter != " "{
		t.Error("Expected ' ' but got" + letter)
	}
}