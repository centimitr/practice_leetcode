package practice_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	a, b := head, head.Next
	for a != b {
		if b.Next == nil || b.Next.Next == nil {
			return false
		}
		a = a.Next
		b = b.Next.Next
	}
	return true
}
