package util

import (
	"io/ioutil"
	"strconv"
)

func ReadInputFromFile(s string) string {
	dat, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func ParseInteger(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return num
}
