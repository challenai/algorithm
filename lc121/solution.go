package solution

func maxProfit(prices []int) int {
	// idea: every time meet an el, if it smaller than minimum till current el, update minimum
	// if this el is a bigger one , just compare the diff with resu
	resu := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i]-min > resu {
			resu = prices[i] - min
		}
	}

	return resu
}
