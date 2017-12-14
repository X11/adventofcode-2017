package main

import (
	"fmt"
	"strconv"

	"github.com/X11/adventofcode-2017/util"
)

func main() {
	input := "xlqgujun"

	grid := [][]int{}

	count := 0

	for y := 0; y < 128; y++ {
		hash := knotHash(fmt.Sprintf("%s-%d", input, y))
		fmt.Printf("hash = %+v\n", hash)
		row := []int{}
		for i, _ := range hash {
			bins := Hex2Bin(hash[i])
			for i, _ := range bins {
				f := 0
				if string(bins[i]) == "1" {
					f = 1
					count++
				}
				row = append(row, f)
			}
		}
		grid = append(grid, row)
	}

	fmt.Printf("grid = %+v\n", grid)
	fmt.Printf("count = %+v\n", count)
}

func Hex2Bin(in byte) string {
	i, _ := strconv.ParseInt(string(in), 16, 0)
	return fmt.Sprintf("%.4b", i)
}

func knotHash(inputStr string) string {
	extraSeq := []int{17, 31, 73, 47, 23}
	input := []int{}
	for i, _ := range inputStr {
		input = append(input, util.ParseInteger(fmt.Sprint(inputStr[i])))
	}
	for _, k := range extraSeq {
		input = append(input, k)
	}

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

	dense := []int{}
	for i := 0; i < 16; i++ {
		offset := i * 16
		s := list[offset]
		for j := 1; j < 16; j++ {
			s = s ^ list[offset+j]
		}
		dense = append(dense, s)
	}

	hash := ""
	for _, s := range dense {
		hash += fmt.Sprintf("%02x", s)
	}

	return hash
}
