package practice_leetcode

func maxArea(height []int) (max int) {
	var h, area int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}
		area = h * (j - i + 1)
		if area > max {
			max = area
		}
	}
	return
}

// max(min(h[i],h[j)*(j-i))
