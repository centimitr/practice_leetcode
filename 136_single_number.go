package practice_leetcode

func singleNumber(nums []int) int {
	a := 0
	for _, n := range nums {
		a = a ^ n
	}
	return a
}
