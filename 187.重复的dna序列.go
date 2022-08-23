/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	ans := make([]string, 0)
	if len(s) < 10 {
		return ans
	}

	record := make(map[string]int, 0)
	for i := 0; i <= len(s)-10; i++ {
		cur := s[i : i+10]
		if num, ok := record[cur]; ok && num == 1 {
			ans = append(ans, cur)
		}
		record[cur]++
	}

	return ans
}

// @lc code=end
