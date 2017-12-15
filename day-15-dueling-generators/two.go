package main

import "fmt"

type Generator struct {
	Value    int
	Factor   int
	Multiply int
}

func (g *Generator) Generate() string {
	for {
		g.Value = (g.Value * g.Factor) % 2147483647
		if g.Value%g.Multiply == 0 {
			bin := fmt.Sprintf("%b", g.Value)
			bin = fmt.Sprintf("%030s", bin)
			return bin[len(bin)-16:]
		}
	}
}

func NewGenerator(v int, f int, m int) *Generator {
	return &Generator{
		Value:    v,
		Factor:   f,
		Multiply: m,
	}
}

func main() {
	a := NewGenerator(703, 16807, 4)
	b := NewGenerator(516, 48271, 8)

	count := 0
	pairs := 5 * 1000 * 1000
	for i := 0; i < pairs; i++ {
		if i%1000000 == 0 {
			fmt.Printf("i = %+v\n", i)
		}
		if a.Generate() == b.Generate() {
			count++
		}
	}
	fmt.Printf("count = %+v\n", count)
}
