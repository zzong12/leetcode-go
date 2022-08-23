/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	ans := len(nums)
	for i := 0; i < len(nums); i++ {
		ans ^= i ^ nums[i]
	}
	return ans
}

// @lc code=end

