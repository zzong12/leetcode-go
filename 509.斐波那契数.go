/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n < 2 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b%1000000007
	}
	return a + b
}

// @lc code=end
