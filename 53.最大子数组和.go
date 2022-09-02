/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 3, -1, -2, -1}
// 	fmt.Println(maxSubArray(nums))
// 	fmt.Println(nums)
// }

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if ans < nums[i] {
			ans = nums[i]
		}
	}
	return ans
}

// @lc code=end
