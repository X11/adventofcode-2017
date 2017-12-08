package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkForDubble(test string) bool {
	words := strings.Split(test, " ")
	for i, word := range words {
		for _, testWord := range words[i+1:] {
			if word == testWord {
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
