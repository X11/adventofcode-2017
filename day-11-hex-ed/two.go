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

func getFurthestDistance(input string) int {
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

	furthest := 0
	for _, step := range steps {
		if step == "" {
			continue
		}
		off := offsets[step]
		pos.X += off.X
		pos.Y += off.Y
		distance := calculateDistance(Offset{0, 0}, pos)
		if distance > furthest {
			furthest = distance
		}
	}

	fmt.Printf("furthest = %+v\n", furthest)
	return furthest
}

func main() {
	input := util.ReadInputFromFile("./input.txt")
	getFurthestDistance(strings.Trim(input, "\n"))
}
