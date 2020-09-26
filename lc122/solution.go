package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	resu := 0
	mini := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] > mini {
			resu += prices[i] - mini
		}
		mini = prices[i]
	}
	return resu
}
