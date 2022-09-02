/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// @lc code=start
func longestPalindrome(s string) string {
	var ans string
	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l = l - 1
			r = r + 1
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l = l - 1
			r = r + 1
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}
	}
	return ans
}

// @lc code=end
