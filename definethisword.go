package main

import (
	"bufio"
	"fmt"
	"github.com/kofoworola/definethephrase/twitter"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(writer,"Server is up and running")

	})
	http.HandleFunc("/twitter/webhook", twitter.CrcCheck)
	if len(os.Getenv("BOT_PORT")) > 1 {
		http.ListenAndServe(":" + os.Getenv("BOT_PORT"), nil)
	} else{
		http.ListenAndServe(":80", nil)
	}
}

func understand(words []string, index int) {
	handle := os.Getenv("HANDLE")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter tweet: ")
	text, _ := reader.ReadString('\n')
	reg := regexp.MustCompile("@" + handle + " (the|this) (word|phrase)")
	indices := reg.FindStringIndex(text)
	if len(indices) != 2 {
		fmt.Println("Could not understand")
		return
	}
	phrase := strings.Trim(text[indices[1]:], " ")
	client := DefinitionClient{Phrase: phrase, Provider: "oxford"}
	client.CheckDefinition()
}
