/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	ans := t[len(t)-1]
	for i := 0; i < len(s); i++ {
		ans ^= s[i] ^ t[i]
	}
	return ans
}

// @lc code=end

