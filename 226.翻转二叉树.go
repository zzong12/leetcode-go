/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertNode(root)
	return root
}

func invertNode(node *TreeNode) {
	node.Left, node.Right = node.Right, node.Left
	if node.Left != nil {
		invertNode(node.Left)
	}
	if node.Right != nil {
		invertNode(node.Right)
	}
}

// @lc code=end
