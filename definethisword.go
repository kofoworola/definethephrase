package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const HANDLE string = "@__define__"

func main() {
	http.HandleFunc("/",twitterWebhook)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func understand(words []string, index int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter tweet: ")
	text, _ := reader.ReadString('\n')
	reg := regexp.MustCompile(HANDLE + " (the|this) (word|phrase)")
	indices := reg.FindStringIndex(text)
	if len(indices) != 2 {
		fmt.Println("Could not understand")
		return
	}
	phrase := strings.Trim(text[indices[1]:]," ")
	client := DefinitionClient{Phrase:phrase,Provider: "oxford"}
	client.CheckDefinition()
}

func twitterWebhook(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"Hello kofo")
}
