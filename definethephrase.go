package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kofoworola/definethephrase/twitter"
	"log"
	"net/http"
	"os"
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
		} else if args[1] == "-delete" && len(args) > 2{
			go twitter.DeleteWebhook(args[2])
		}
	}
	setUpServer()
}

func setUpServer() {
	//Register routes with mux
	fmt.Println("Starting Server")
	m := mux.NewRouter()
	m.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Server is up and running")
	})
	m.HandleFunc("/twitter/webhook", CrcCheck).Methods("GET")
	m.HandleFunc("/twitter/webhook", WebhookHandler).Methods("POST")
	server := &http.Server{
		Handler: m,
	}
	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	server.Addr = ":" + port
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("Error hit on starting server")
		fmt.Println(err.Error())
	}
}