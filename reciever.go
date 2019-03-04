package main

import "strings"

func turnToWords(tweet string) []string{
	return strings.Split(tweet," ");
}
