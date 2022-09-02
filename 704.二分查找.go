/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start

func search(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (r + l) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// @lc code=end
