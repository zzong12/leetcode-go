/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

func lengthOfLongestSubstring(s string) int {

	bs := []byte(s)
	m := make(map[byte]int)
	l, max, size := 0, 0, len(bs)

	for i := 0; i < size; i++ {
		v := bs[i]
		if j, ok := m[v]; ok {
			if i-l > max {
				max = i - l
			}
			if j+1 > l {
				l = j + 1
			}
		}
		m[v] = i
	}

	if size-l > max {
		max = size - l
	}

	return max
}

// @lc code=end
