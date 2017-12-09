package main

import (
	"fmt"
	"github.com/X11/adventofcode-2017/util"
	"strings"
)

func main() {
	input := util.ReadInputFromFile("./input.txt")

	fmt.Printf("Count(input) = %+v\n", Count(input))
}

func Count(i string) int {
	chars := strings.Split(i, "")

	count := 0
	depth := 0
	garbage := false
	ignore := false

	fmt.Printf("i = %+v\n", i)

	for _, c := range chars {
		if ignore {
			ignore = false
			continue
		}
		if garbage {
			switch c {
			case ">":
				garbage = false
				break
			case "!":
				ignore = true
				break
			}
			continue
		}
		switch c {
		case "{":
			depth++
			break
		case "}":
			count += depth
			depth--
			break
		case "!":
			ignore = true
			break
		case "<":
			garbage = true
			break
		}
	}

	return count
}
