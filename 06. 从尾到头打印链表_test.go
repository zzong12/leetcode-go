package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	} else if head.Next == nil {
		return []int{head.Val}
	}
	// cur := head
	// var ans []int
	// for cur != nil {
	// 	tmp := []int{cur.Val}
	// 	ans = append(tmp, ans...)
	// 	cur = cur.Next
	// }

	cur := head
	var ans []int
	for cur != nil {
		ans = append(ans, cur.Val)
		cur = cur.Next
	}

	size := len(ans)
	half := size / 2
	for i := 0; i < half; i++ {
		ans[i], ans[size-1-i] = ans[size-1-i], ans[i]
	}

	return ans
}
