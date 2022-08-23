package main

import "fmt"

func Test_reverseLeftWords() {
	fmt.Sprintln(reverseLeftWords("abcdefg", 2))
}

func reverseLeftWords(s string, n int) string {
	if n == len(s) {
		return s
	}
	ans := []byte(s)
	ans = append(ans, s[0:n]...)

	return string(ans[n:])
}
