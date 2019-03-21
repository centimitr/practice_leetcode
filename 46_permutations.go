package practice_leetcode

// via swap
func fn(nums []int, start int, res *[][]int) {
	if len(nums) == start {
		p := make([]int, len(nums))
		copy(p, nums)
		*res = append(*res, p)
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		fn(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permute(nums []int) (res [][]int) {
	fn(nums, 0, &res)
	return
}

// via recursion + copy
func permute1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var s [][]int
	for i, n := range nums {
		ns := make([]int, len(nums)-1)
		copy(ns[:], nums[:i])
		copy(ns[i:], nums[i+1:])
		for _, v := range permute(ns) {
			p := make([]int, len(nums))
			p[0] = n
			copy(p[1:], v)
			s = append(s, p)
		}
		_ = n
	}
	return s
}

// python
//from itertools import permutations
//class Solution:
//def permute(self, nums: List[int]) -> List[List[int]]:
//return list(permutations(nums, len(nums)))
