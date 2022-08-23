/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	// 0,1,1,2,3,5,8,13...
	// 1,2,3,4,5,6,

	// 1)
	// if n < 2 {
	// 	return n
	// }
	// return fib(n-2) + fib(n-1)

	// 2)
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return a + b
}

// @lc code=end

