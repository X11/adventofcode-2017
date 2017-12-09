package main

import "testing"

var fibTests = []struct {
	n        string // input
	expected int    // expected result
}{
	{"{}", 1},
	{"{{{}}}", 6},
	{"{{},{}}", 5},
	{"{{{},{},{{}}}}", 16},
	{"{<a>,<a>,<a>,<a>}", 1},
	{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
	{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
	{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
}

func TestCount(t *testing.T) {
	for _, tt := range fibTests {
		actual := Count(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
