/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// package main

// import "fmt"

// func main() {
// 	matrix := [][]int{
// 		// {1, 2, 3, 4, 5},
// 		// {6, 7, 8, 9, 10},
// 		// {11, 12, 13, 14, 15},
// 		// {16, 17, 18, 19, 20},
// 		// {21, 22, 23, 24, 25},
// 	}
// 	target := 19
// 	fmt.Println(searchMatrix(matrix, target))
// }

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	xLen := len(matrix[0])

	firstLine := matrix[0]
	l, r := 0, xLen
	for l < r {
		mid := l + (r-l)>>1
		if firstLine[mid] == target {
			return true
		} else if firstLine[mid] > target {
			r = mid
		} else if firstLine[mid] < target {
			l = mid + 1
		}
	}

	if l == 0 && firstLine[0] > target {
		return false
	}

	endColumn := l
	for i := 0; i < endColumn; i++ {
		l, r = 0, len(matrix)
		for l < r {
			mid := l + (r-l)>>1
			if matrix[mid][i] == target {
				return true
			} else if matrix[mid][i] > target {
				r = mid
			} else if matrix[mid][i] < target {
				l = mid + 1
			}
		}
	}

	return false
}

// @lc code=end
