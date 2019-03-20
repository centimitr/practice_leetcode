package practice_leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1
	for ; i >= 0; i-- {
		if j >= 0 && (k < 0 || nums1[j] > nums2[k]) {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
}

// as the two slices are sorted, merge from the end can avoid moving
