package practice_leetcode

import (
	"container/heap"
)

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() (v interface{}) {
	v = (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return
}

func findKthLargest(nums []int, k int) int {
	h := intHeap(nums)
	heap.Init(&h)
	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}
