/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start

func rotate_copyarray(nums []int, k int) {
	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r[(i+k)%len(nums)] = nums[i]
	}
	for i := 0; i < len(r); i++ {
		nums[i] = r[i]
	}
}

func rotate_reverse(nums []int) {
	for i, n := 0, len(nums); i < n>>1; i++ {
		j := n - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate(nums []int, k int) {
	p := k % len(nums)
	rotate_reverse(nums)
	rotate_reverse(nums[:p])
	rotate_reverse(nums[p:])
}

// @lc code=end
