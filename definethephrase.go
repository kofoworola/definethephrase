package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kofoworola/definethephrase/twitter"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}
	if args := os.Args; len(args) > 1{
		if args[1] == "-gethook"{
			go twitter.GetWebhook()
		}else if args[1] == "-subscribe"{
			go twitter.SubscribeWebhook()
		}
	}
	setUpServer()
}

func setUpServer() {
	//Register routes with mux
	m := mux.NewRouter()
	m.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Server is up and running")
	})
	m.HandleFunc("/twitter/webhook", twitter.CrcCheck).Methods("GET")
	m.HandleFunc("/twitter/webhook", twitter.WebhookHandler).Methods("POST")
	server := &http.Server{
		Handler: m,
	}
	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	server.Addr = ":" + port
	server.ListenAndServe()
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
