package main

import (
	"fmt"
	"regexp"
	"strings"
)

const HANDLE string = "@__define__"

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter tweet: ")
	//text, _ := reader.ReadString('\n')
	text := "@__define__ the word fool"
	reg := regexp.MustCompile(HANDLE + " (the|this) (word|phrase)")
	indices := reg.FindStringIndex(text)
	if len(indices) != 2 {
		fmt.Println("Could not understand")
		return
	}
	phrase := strings.Trim(text[indices[1]:]," ")
	client := DefinitionClient{Phrase:phrase,Provider: "oxford"}
	client.CheckDefinition()
	fmt.Println(phrase)
}

func understand(words []string, index int) {
	phrase := words[index]
	fmt.Println(phrase)
}
