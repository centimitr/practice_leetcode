package practice_leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	j := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
