package main

import (
	"fmt"
	"strings"

	"github.com/X11/adventofcode-2017/util"
)

type Offset struct {
	X int
	Y int
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func calculateDistance(a Offset, b Offset) int {
	return (abs(a.X-b.X) + abs(a.X+a.Y-b.X-b.Y) + abs(a.Y-b.Y)) / 2
}

func getDistance(input string) int {
	steps := strings.Split(input, ",")

	offsets := map[string]Offset{
		"nw": Offset{-1, 0},
		"n":  Offset{0, -1},
		"ne": Offset{1, -1},
		"se": Offset{1, 0},
		"s":  Offset{0, 1},
		"sw": Offset{-1, 1},
	}

	pos := Offset{0, 0}

	for _, step := range steps {
		if step == "" {
			continue
		}

		fmt.Printf("step = %+v\n", step)
		off := offsets[step]
		fmt.Printf("pos = %+v\n", pos)
		fmt.Printf("off = %+v\n", off)

		pos.X += off.X
		pos.Y += off.Y
		fmt.Printf("pos = %+v\n", pos)
	}

	distance := calculateDistance(Offset{0, 0}, pos)

	fmt.Printf("distance = %+v\n", distance)
	return distance
}

func main() {
	var fibTests = []struct {
		n        string // input
		expected int    // expected result
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}

	for _, tt := range fibTests {
		actual := getDistance(tt.n)
		if actual != tt.expected {
			fmt.Printf("getDistance(%s): expected %d, actual %d\n", tt.n, tt.expected, actual)
		}
	}

	input := util.ReadInputFromFile("./input.txt")
	getDistance(strings.Trim(input, "\n"))
}
