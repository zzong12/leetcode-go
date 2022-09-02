/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	headNew := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		nodeNew := cur.Next
		cur.Next = cur.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew
}

// @lc code=end
