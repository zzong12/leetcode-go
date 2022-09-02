package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var ans [][]int

	for len(queue) != 0 {
		rlen := len(queue)
		var line []int
		for i := 0; i < rlen; i++ {
			node := queue[i]
			line = append(line, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[rlen:]
		if len(ans)&1 == 1 {
			lineLen := len(line)
			for j := 0; j < len(line)>>1; j++ {
				line[j], line[lineLen-1-j] = line[lineLen-1-j], line[j]
			}
		}
		ans = append(ans, line)
	}
	return ans
}
