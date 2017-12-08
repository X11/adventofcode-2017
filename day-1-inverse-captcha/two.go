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

func getComparismentIndex(curPos int, offset int, length int) int {
	offsetted := curPos + offset
	if offsetted >= length {
		return offsetted - length
	}

	return offsetted
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter in the puzzle")
	}

	puzzle := os.Args[1]
	offset := len(puzzle) / 2
	len := len(puzzle)
	bytes := []byte(puzzle)

	fmt.Printf("puzzle = %+v\n", puzzle)
	fmt.Printf("offset = %+v\n", offset)

	sum := 0
	for index, b := range bytes {
		if b == bytes[getComparismentIndex(index, offset, len)] {
			sum += getValue(puzzle[index])
		}
	}

	fmt.Printf("sum = %+v\n", sum)
}
