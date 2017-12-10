package main

import "fmt"

func main() {

	input := []int{94, 84, 0, 79, 2, 27, 81, 1, 123, 93, 218, 23, 103, 255, 254, 243}

	list := []int{}
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	listLen := len(list)

	curPos := 0
	skipSize := 0

	setCurPos := func(skip int) {
		curPos += skip
		if curPos >= listLen {
			curPos -= listLen
		}
	}

	getPos := func(offset int) int {
		if curPos+offset >= listLen {
			return -listLen + curPos + offset
		}
		return curPos + offset
	}

	for _, length := range input {
		fmt.Printf("list = %+v\n", list)
		fmt.Printf("curPos = %+v\n", curPos)
		fmt.Printf("length = %+v\n", length)
		for i := 0; i < length/2; i++ {
			// We swapped all
			if i > length/2 {
				break
			}

			a := getPos(i)
			b := getPos(length - i - 1)
			fmt.Printf("a) %+v\tb) %+v\n", a, b)
			swap := list[a]
			list[a] = list[b]
			list[b] = swap
		}
		fmt.Printf("list = %+v\n", list)
		fmt.Printf("skipSize = %+v\n", skipSize)
		setCurPos(length + skipSize)
		skipSize++
	}

	fmt.Printf("list = %+v\n", list)
	fmt.Printf("list[0]*list[1] = %+v\n", list[0]*list[1])
}
