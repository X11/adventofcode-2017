package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func sortWord(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func checkForDubble(test string) bool {
	words := strings.Split(test, " ")
	for i, word := range words {
		sorted := sortWord(word)
		for _, testWord := range words[i+1:] {
			if sorted == sortWord(testWord) {
				return true
			}
		}
	}

	return false
}

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	phrases := strings.Split(string(dat), "\n")
	for _, phrase := range phrases {
		if phrase == "" {
			continue
		}

		if !checkForDubble(phrase) {
			sum++
		}
	}

	fmt.Printf("sum = %+v\n", sum)
}
