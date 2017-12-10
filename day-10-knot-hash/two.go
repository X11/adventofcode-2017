package main

import (
	"fmt"

	"github.com/X11/adventofcode-2017/util"
)

func main() {

	inputStr := "94,84,0,79,2,27,81,1,123,93,218,23,103,255,254,243"
	extraSeq := []int{17, 31, 73, 47, 23}
	input := []int{}
	for i, _ := range inputStr {
		input = append(input, util.ParseInteger(fmt.Sprint(inputStr[i])))
	}
	for _, k := range extraSeq {
		input = append(input, k)
	}

	fmt.Printf("input = %+v\n", input)

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
		pos := curPos + offset
		for pos >= listLen {
			pos = -listLen + pos
		}
		return pos
	}

	for r := 0; r < 64; r++ {
		for _, length := range input {
			for i := 0; i < length/2; i++ {
				// We swapped all
				if i > length/2 {
					break
				}

				a := getPos(i)
				b := getPos(length - i - 1)
				swap := list[a]
				list[a] = list[b]
				list[b] = swap
			}
			setCurPos(length + skipSize)
			skipSize++
		}
	}

	fmt.Printf("list = %+v\n", list)

	dense := []int{}
	for i := 0; i < 16; i++ {
		offset := i * 16
		s := list[offset]
		for j := 1; j < 16; j++ {
			s = s ^ list[offset+j]
		}
		dense = append(dense, s)
	}

	fmt.Printf("dense = %+v\n", dense)

	hash := ""
	for _, s := range dense {
		hash += fmt.Sprintf("%02x", s)
	}
	fmt.Printf("hash = %+v\n", hash)
}
