/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r, s := head, head, 0
	for r.Next != nil {
		r = r.Next
		if s == n {
			l = l.Next
		} else {
			s++
		}
	}
	if s < n {
		return head.Next
	}
	if l.Next == nil {
		return nil
	}
	l.Next = l.Next.Next
	return head
}

// @lc code=end
