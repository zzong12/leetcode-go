/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		v := numbers[l] + numbers[r]
		if v == target {
			return []int{l + 1, r + 1}
		} else if v > target {
			r--
		} else {
			l++
		}
	}
	return nil
}

// @lc code=end
