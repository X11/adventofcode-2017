package main

import "fmt"

func ulam(n int, b int) int {
	return 4*(n*n) + (b * n) + 1
}

func main() {
	input := 289326
	fmt.Printf("input = %+v\n", input)

	order := []int{-3, -2, -1, 0, 1, 2, 3, 4}

main:
	for i := 0; true; i++ {
		for h, o := range order {
			if ulam(i, o) > input {
				var r int
				if h%2 == 1 {
					r = input - ulam(i, o-1)
				} else {
					r = ulam(i, o) - input
				}
				fmt.Printf("i = %+v\n", i)
				fmt.Printf("o = %+v\n", o)
				fmt.Printf("r + i = %+v\n", r+i)
				break main
			}
		}
	}

}
