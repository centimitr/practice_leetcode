package practice_leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	head := new(ListNode)
	n := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			n.Next = l1
			l1 = l1.Next
		}else {
			n.Next = l2
			l2 = l2.Next
		}
		n = n.Next
	}
	if l1 == nil {
		n.Next = l2
	}else if l2 == nil {
		n.Next = l1
	}
	return head.Next
}