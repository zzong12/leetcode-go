package main

/**
 * https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil &&
		B != nil &&
		(recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}
