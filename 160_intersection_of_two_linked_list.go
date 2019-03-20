package practice_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	n1, n2 := headA, headB
	cnt := 0
	for {
		if n1 == n2 {
			return n1
		}
		n1 = n1.Next
		n2 = n2.Next
		if n1 == nil {
			n1 = headB
			cnt++
		}
		if n2 == nil {
			n2 = headA
			cnt++
		}
		if cnt > 2 {
			return nil
		}
	}
}

// the key to solve this problem is, the length of listA and listB is not the same
// so we can instead check listA+listB and listB+listA
