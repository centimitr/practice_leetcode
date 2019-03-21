package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{
		3,
		&TreeNode{1, nil,
			&TreeNode{2, nil, nil}},
		&TreeNode{4, nil, nil},
	}
	fmt.Println(kthSmallest(r, 1))
}

type visitFn func(v int) (stop bool)

func inOrder(node *TreeNode, visit visitFn) (stop bool) {
	if node == nil {
		return false
	}
	if inOrder(node.Left, visit) || visit(node.Val) || inOrder(node.Right, visit) {
		return true
	}
	return false
}

func kthSmallest(root *TreeNode, k int) int {
	var val int
	inOrder(root, func(v int) bool {
		if k > 1 {
			k--
			return false
		}
		val = v
		return true
	})
	return val
}
