/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]w
			i++
		}
		j++
	}
}

// @lc code=end
