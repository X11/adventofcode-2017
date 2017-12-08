package main

import (
	"fmt"
	"os"
	"strconv"
)

func getValue(r byte) int {
	number, _ := strconv.Atoi(string(r))
	return number
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter in the puzzle")
	}

	puzzle := os.Args[1]

	fmt.Printf("puzzle = %+v\n", puzzle)

	sum := 0
	var prev rune
	for index, runeValue := range puzzle {
		if runeValue == prev {
			sum += getValue(puzzle[index])
		}
		prev = runeValue
	}

	if puzzle[0] == puzzle[len(puzzle)-1] {
		sum += getValue(puzzle[0])
	}

	fmt.Printf("sum = %+v\n", sum)
}
