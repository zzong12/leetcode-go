/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	size := len(nums)
	l, r := 0, size
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			r = mid
		} else if nums[mid] > target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	if l == size || nums[l] != target {
		return []int{-1, -1}
	}

	for r = l; r < size && nums[r] == target; r++ {

	}

	return []int{l, r - 1}
}

// @lc code=end
