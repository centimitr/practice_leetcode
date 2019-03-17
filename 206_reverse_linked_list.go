package practice_leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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
