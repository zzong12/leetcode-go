package main

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxValue(grid [][]int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				a, b := grid[i][j-1], grid[i-1][j]
				if a > b {
					grid[i][j] += a
				} else {
					grid[i][j] += b
				}
			}
		}
	}
	return grid[m-1][n-1]
}
