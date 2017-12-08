package main

import "fmt"

func getPrint(a []int) string {
	return fmt.Sprint(a)
}

func findHighestIndex(a []int) int {
	var max int = a[0]
	var i int = 0
	for j, value := range a {
		if max < value {
			max = value
			i = j
		}
	}
	return i
}

func main() {
	prints := make(map[string]bool)
	banks := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
	bl := len(banks)

	getIndex := func(start int, offset int) int {
		offset = offset % bl
		if start+offset >= bl {
			return start + offset - bl
		}
		return start + offset
	}

	steps := 0
	for {
		steps++

		highestIndex := findHighestIndex(banks)
		blocks := banks[highestIndex]
		banks[highestIndex] = 0
		offset := 0
		for blocks > 0 {
			offset++
			banks[getIndex(highestIndex, offset)]++
			blocks--
		}

		print := getPrint(banks)
		_, ok := prints[print]
		if ok {
			fmt.Printf("steps = %+v\n", steps)
			break
		} else {
			prints[print] = true
		}
	}
}
