/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords_single(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	b := []byte(s)
	for l, r := 0, 1; r < len(b); {
		if b[r] == ' ' {
			reverseWords_single(b[l:r])
			l = r + 1
		} else if r == len(b)-1 {
			reverseWords_single(b[l:])
		}
		r++
	}
	return string(b)
}

// @lc code=end
