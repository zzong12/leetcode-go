/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	rnums := make([]int, len(nums))
	for i := r; i >= 0; i-- {
		if v, w := nums[l]*nums[l], nums[r]*nums[r]; v > w {
			rnums[i] = v
			l++
		} else {
			rnums[i] = w
			r--
		}
	}
	return rnums
}

// @lc code=end
