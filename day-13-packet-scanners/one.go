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

func (f *Firewall) Check() bool {
	if f.Len == 0 && f.Cur == 0 {
		return false
	}

	return f.Cur == 0
}

func (f *Firewall) Step() {
	if f.Len > 0 {
		if f.Cur == f.Len-1 || f.Cur == 0 {
			f.Dir = !f.Dir
		}

		switch f.Dir {
		case true:
			f.Cur++
			break
		case false:
			f.Cur--
			break
		}
	}
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	firewalls := []*Firewall{}

	steps := strings.Split(input, "\n")
	steps = steps[:len(steps)-1]
	max := util.ParseInteger(strings.Split(steps[len(steps)-1], ": ")[0])
	for i := 0; i <= max; i++ {
		firewalls = append(firewalls, &Firewall{0, 0, false})
	}

	for _, step := range steps {
		parts := strings.Split(step, ": ")
		i := util.ParseInteger(parts[0])
		c := util.ParseInteger(parts[1])
		firewalls[i].Len = c
	}

	severity := 0
	for i := 0; i <= max; i++ {
		if firewalls[i].Check() {
			severity += i * firewalls[i].Len
		}

		for k := 0; k <= max; k++ {
			firewalls[k].Step()
		}
	}

	fmt.Printf("severity = %+v\n", severity)
}
