package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	instructionStrings := strings.Split(string(dat), "\n")
	instruction := []int{}
	for _, i := range instructionStrings {
		if i == "" {
			continue
		}
		v, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		instruction = append(instruction, v)
	}

	pos := 0
	step := 0
	for pos < len(instruction) {
		jump := instruction[pos]
		if jump >= 3 {
			instruction[pos] += -1
		} else {
			instruction[pos] += 1
		}
		pos += jump
		step++
	}
	fmt.Printf("step = %+v\n", step)
}
