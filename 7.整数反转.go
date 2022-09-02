/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
// @lc code=start

func reverse(x int) int {
	r := 0
	for x != 0 {
		r = x%10 + r*10
		x /= 10
	}
	if r < math.MinInt32 || r > math.MaxInt32 {
		return 0
	}
	return r
}

// @lc code=end

