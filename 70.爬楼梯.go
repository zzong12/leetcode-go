/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, (a + b)
	}

	return (a + b)
}

// @lc code=end
