package practice_leetcode

func reverseList(head *ListNode) *ListNode {
	n := head
	var prev *ListNode
	var next *ListNode
	for n != nil {
		next = n.Next
		n.Next = prev
		prev = n
		n = next
	}
	return prev
}
