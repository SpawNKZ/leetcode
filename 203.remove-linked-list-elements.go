package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: -1}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev.Next = curr
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}
