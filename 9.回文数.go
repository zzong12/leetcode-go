/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	helf := len(strX) / 2
	for i := 0; i < helf; i++ {
		if strX[i] != strX[len(strX)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end

