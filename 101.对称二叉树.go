/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricNode(root.Left, root.Right)
}

func isSymmetricNode(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		return false
	}
	return a.Val == b.Val &&
		isSymmetricNode(a.Left, b.Right) &&
		isSymmetricNode(a.Right, b.Left)

}

// @lc code=end
