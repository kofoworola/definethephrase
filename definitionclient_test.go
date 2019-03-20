package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestOxfordCheck(t *testing.T){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}
	definitionClient := DefinitionClient{"ace","oxford",""}
	definition := definitionClient.CheckDefinition()
	fmt.Println(definition)
	if len(definition) < 1{
		t.Error("Was expecting definitions but got "+ definition)
	}
}
