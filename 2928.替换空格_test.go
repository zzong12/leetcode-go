// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// 输入：s = "We are happy."
// 输出："We%20are%20happy."

// 0 <= s 的长度 <= 10000

package main

import (
	"fmt"
	"unsafe"
)

func Test_replaceSpace() {
	fmt.Sprintln(replaceSpace("23313"))
}

var replaceStr = []byte("%20")

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	var ans []byte
	for _, b := range []byte(s) {
		if b == ' ' {
			ans = append(ans, replaceStr...)
		} else {
			ans = append(ans, b)
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}
