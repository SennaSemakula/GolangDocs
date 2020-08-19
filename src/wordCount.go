package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

type Word struct {
	word string
	count int
}

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	for _, v := range strings.Fields(s){
		_, ok := wordMap[v]
		if ok {
			wordMap[v] += 1
		} else {
			wordMap[v] = 1
		}
	}
	fmt.Println(wordMap)
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
