package main

import (
	"strconv"
)

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

f(−1)=0
f(0)=1
f(1)=f(0) + f(-1) = 1
f(i)=f(i−1)+f(i−2)[i−1≥0,10≤x≤25]
*/

// func main() {
// 	fmt.Println(translateNum(12258))
// }

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	src := strconv.Itoa(num)

	a, b, val := 1, 1, src[0:2]
	if val >= "10" && val <= "25" {
		b = 2
	}
	for i := 2; i < len(src); i++ {
		val := src[i-1 : i+1]
		tmp := b
		if val >= "10" && val <= "25" {
			tmp += a
		}
		a, b = b, tmp
	}

	return b
}
