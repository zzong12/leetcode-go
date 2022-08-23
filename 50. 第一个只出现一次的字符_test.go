package main

import (
	"fmt"
)

func Test_firstUniqChar() {
	s := "abaccdeff"
	r := findRepeatNumber(s)
	fmt.Sprintln(r)
}

func firstUniqChar(s string) byte {
	dict := [26]int{}
	bs := []byte(s)
	for _, b := range bs {
		idx := b - 'a'
		dict[idx]++
	}
	for _, b := range bs {
		idx := b - 'a'
		if dict[idx] == 1 {
			return b
		}
	}
	return ' '
}
