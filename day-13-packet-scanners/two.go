package main

import (
	"fmt"
	"strings"

	"github.com/X11/adventofcode-2017/util"
)

type Firewall struct {
	Len int
	Cur int
	Dir bool
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	firewalls := make(map[int]int)

	steps := strings.Split(input, "\n")
	steps = steps[:len(steps)-1]
	for _, step := range steps {
		parts := strings.Split(step, ": ")
		i := util.ParseInteger(parts[0])
		c := util.ParseInteger(parts[1])
		firewalls[i] = c
	}

	caught := true
	delay := 0
	for caught == true {
		caught = false
		for k, v := range firewalls {
			if (k+delay)%(2*(v-1)) == 0 {
				delay++
				caught = true
				break
			}
		}
		fmt.Printf("delay = %+v\n", delay)
	}
	fmt.Printf("delay = %+v\n", delay)
}
