package main

import (
	"math"
	"sort"
)

func intsSum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var minSum int
	var minDiff float64
	for p1, p2 := 0, 2; p2 != len(nums); p1, p2 = p1+1, p2+1 {
		sum := intsSum(nums[p1 : p2+1])
		diff := math.Abs(float64(target - sum))
		if p1 == 0 || diff < minDiff {
			minDiff = diff
			minSum = sum
		}
	}
	return minSum
}

func main() {
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
}
