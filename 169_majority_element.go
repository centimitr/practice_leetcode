package practice_leetcode

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	max := 0
	maxN := 0
	for n, cnt := range m {
		if cnt > max {
			max = cnt
			maxN = n
		}
	}
	return maxN
}

// or sort and get the center element
