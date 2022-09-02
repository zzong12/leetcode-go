/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

package main

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	dict := make([]int, 128)
	for i := 0; i < 128; i++ {
		dict[i] = -1
	}

	i, ans := -1, 0
	for j, b := range []byte(s) {
		if idx := dict[b]; idx >= 0 {
			i = idx
		}
		dict[b] = j
		size := j - i
		if size > ans {
			ans = size
		}
	}
	return ans
}

// @lc code=end
