package practice_leetcode

type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
}

func (ms *MinStack) Pop() {
	if len(ms.stack) != 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) != 0 {
		return ms.stack[len(ms.stack)-1]
	}
	return 0
}

func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		return 0
	}
	min := ms.stack[0]
	for i := 1; i < len(ms.stack); i++ {
		if ms.stack[i] < min {
			min = ms.stack[i];
		}
	}
	return min
}
