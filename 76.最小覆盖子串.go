/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	if s == t {
		return s
	}

	tc, wc := [64]int{}, [64]int{}
	for i := 0; i < len(t); i++ {
		tc[t[i]-'A']++
	}

	var ans string
	for l, r, cnt := 0, 0, 0; r < len(s); r++ {
		rv := s[r] - 'A'

		wc[rv]++
		if wc[rv] <= tc[rv] {
			cnt++
		}

		for ; l < r && wc[s[l]-'A'] > tc[s[l]-'A']; l++ {
			wc[s[l]-'A']--
		}

		if cnt == len(t) && (ans == "" || len(ans) > r-l+1) {
			ans = s[l : r+1]
		}

	}

	return ans
}

// @lc code=end
