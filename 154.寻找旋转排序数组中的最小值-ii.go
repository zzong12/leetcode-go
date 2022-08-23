/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// package main

// import "fmt"

// func main() {
// 	r := findMin([]int{1, 3, 5})
// 	fmt.Sprintln(r)
// }

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] == nums[r] {
			r--
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		}
	}
	return nums[l]
}

// @lc code=end
