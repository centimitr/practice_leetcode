package practice_leetcode

func subsets(nums []int) [][]int {
	sets := [][]int{{}}
	for _, n := range nums {
		for _, set := range sets {
			tmp := make([]int, len(set)+1)
			copy(tmp, set)
			tmp[len(set)] = n
			sets = append(sets, tmp)
		}
	}
	return sets
}

// think iterating over each number
