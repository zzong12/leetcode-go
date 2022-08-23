/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	bitWords := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			bitWords[i] |= 1 << (words[i][j] - 'a')
		}
	}

}

// @lc code=end

