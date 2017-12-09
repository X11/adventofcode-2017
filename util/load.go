package util

import "io/ioutil"

func ReadInputFromFile(s string) string {
	dat, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	return string(dat)
}
