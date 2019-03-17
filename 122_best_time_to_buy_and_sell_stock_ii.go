package practice_leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	pos := 0
	profit := 0
	for i := 1; i < len(prices); i++ {
		cur := prices[i]
		start := prices[pos]
		if cur > start {
			profit += cur - start
		}
		pos = i
	}
	return profit
}
