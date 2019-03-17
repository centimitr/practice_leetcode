package practice_leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
		return
	}
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
	next.Next = nil
}
