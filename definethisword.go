package main

import (
	"bufio"
	"fmt"
	"os"
)

const HANDLE string = "@_meaningofthephrase"

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter tweet: ")
	text, _ := reader.ReadString('\n')
	index := -1
	words := turnToWords(text);
	for i,word := range words {
		if word == HANDLE{
			index = i
		}
	}
	if index == -1{
		fmt.Println("Can not understand")
		return;
	}
	understand(words,index);
}

func understand(words []string, index int){
	phrase := words[index]
}

