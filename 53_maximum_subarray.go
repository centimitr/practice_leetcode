package practice_leetcode

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, n := range nums {
		// sum>0, sum+n>n => sum+n
		// sum<=0, sum+n<n => n
		if sum > 0 {
			sum += n
		} else {
			sum = n
		}
		// compare current value with former max
		if sum > max {
			max = sum
		}
	}
	return max
}

// f(n) = max(f(n-1)+nums[n], nums[n])
