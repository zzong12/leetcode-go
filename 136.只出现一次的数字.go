/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

// @lc code=end

