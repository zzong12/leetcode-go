/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r]%2 == 0 { // 偶数
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
	return nums
}

// @lc code=end

