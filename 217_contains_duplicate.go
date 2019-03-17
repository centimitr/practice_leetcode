package practice_leetcode

func containsDuplicate(nums []int) bool {
	var cur int
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i]
		for j := i + 1; j < len(nums); j++ {
			if cur == nums[j] {
				return true
			}
		}
	}
	return false
}
