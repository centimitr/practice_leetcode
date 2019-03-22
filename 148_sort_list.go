package practice_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := split(head)
	return merge(sortList(mid), sortList(head))
}

func split(head *ListNode) (mid *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid = slow.Next
	slow.Next = nil
	return
}

func merge(h1, h2 *ListNode) *ListNode {
	head := &ListNode{}
	n := head
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			n.Next = h1
			h1 = h1.Next
		} else {
			n.Next = h2
			h2 = h2.Next
		}
		n = n.Next
	}
	if h1 != nil {
		n.Next = h1
	} else if h2 != nil {
		n.Next = h2
	}
	return head.Next
}
