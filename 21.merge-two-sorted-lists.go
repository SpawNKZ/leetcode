package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	return res
}
