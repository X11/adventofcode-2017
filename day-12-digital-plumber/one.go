package main

import (
	"fmt"
	"strings"

	"github.com/X11/adventofcode-2017/util"
)

type Pipe struct {
	ID            int
	Connected     []*Pipe
	SelfReference bool
}

func (p *Pipe) AddConnected(pipe *Pipe) {
	p.Connected = append(p.Connected, pipe)
}

func (p *Pipe) RemoveConnectedDuplicates() {
	encountered := map[*Pipe]bool{}
	result := []*Pipe{}
	for _, cp := range p.Connected {
		if !encountered[cp] == true {
			encountered[cp] = true
			result = append(result, cp)
		}
	}
	p.Connected = result
}

func getFirstKeyForIntPipeMap(m map[int]*Pipe) int {
	for k := range m {
		return k
	}
	return -1
}

func CountConnectedRecursive(p *Pipe, counted map[int]bool) int {
	counted[p.ID] = true
	c := 1
	for _, cp := range p.Connected {
		if counted[cp.ID] != true {
			c += CountConnectedRecursive(cp, counted)
		}
	}

	return c
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	pipes := make(map[int]*Pipe)

	findOrCreatePipe := func(id int) *Pipe {
		_, ok := pipes[id]
		if !ok {
			pipes[id] = &Pipe{
				ID: id,
			}
		}
		return pipes[id]
	}

	for _, p := range strings.Split(input, "\n") {
		if p == "" {
			continue
		}

		parts := strings.Split(p, " <-> ")
		fmt.Printf("parts = %+v\n", parts)

		id := util.ParseInteger(parts[0])
		pipe := findOrCreatePipe(id)
		for _, c := range strings.Split(parts[1], ",") {
			if c == "" {
				continue
			}

			cid := util.ParseInteger(strings.TrimSpace(c))
			if id == cid {
				pipe.SelfReference = true
				continue
			}

			cpipe := findOrCreatePipe(cid)
			cpipe.AddConnected(pipe)
			pipe.AddConnected(cpipe)
		}

	}

	for _, v := range pipes {
		v.RemoveConnectedDuplicates()
	}

	groups := 0
	for {
		first := getFirstKeyForIntPipeMap(pipes)
		if first < 0 {
			break
		}

		counted := map[int]bool{}
		CountConnectedRecursive(pipes[first], counted)
		for i := range counted {
			delete(pipes, i)
		}
		fmt.Printf("len(counted) = %+v\n", len(counted))
		groups++
	}
	fmt.Printf("groups = %+v\n", groups)

}
