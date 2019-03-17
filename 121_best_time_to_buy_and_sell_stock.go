package practice_leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := 0 - min
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			p := prices[i] - min
			if p > profit {
				profit = p
			}
		}
	}
	if profit > 0 {
		return profit
	}
	return 0
}
