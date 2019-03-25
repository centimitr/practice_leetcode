package practice_leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	n := l
	for n != nil {
		fmt.Print(n.Val)
		fmt.Print("->")
		n = n.Next
	}
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
