/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)

	if len1 > len2 {
		return false
	}

	c1, c2 := [26]int{}, [26]int{}

	for i := 0; i < len1; i++ {
		c1[s1[i]-'a']++
		c2[s2[i]-'a']++
	}

	if c1 == c2 {
		return true
	}

	for r := len1; r < len2; r++ {
		l := r - len1
		c2[s2[r]-'a']++
		c2[s2[l]-'a']--
		if c1 == c2 {
			return true
		}
	}

	return false

}

// @lc code=end
