/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// @lc code=end
