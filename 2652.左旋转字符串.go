package main

func reverseLeftWords(s string, n int) string {
	if n == len(s) {
		return s
	}
	ans := []byte(s)
	ans = append(ans, s[0:n]...)

	return string(ans[n:])
}
