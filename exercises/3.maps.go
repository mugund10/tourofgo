package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordcount := make(map[string]int)
	for _ , val := range words {
		wordcount[val]++
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}
