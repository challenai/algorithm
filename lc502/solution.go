package solution

func findMaximizedCapital(k, w int, Profits, Captical []int) int {
	return w + loop(k, w, Profits, Captical)
}

func loop(k, w int, Profits, Captical []int) int {
	if k == 0 || len(Profits) != len(Captical) || len(Profits) == 0 {
		return 0
	}
	currentMaxProfitIdx := -1
	for i := 0; i < len(Captical); i++ {
		if w >= Captical[i] && (currentMaxProfitIdx == -1 || Profits[i] > Profits[currentMaxProfitIdx]) {
			currentMaxProfitIdx = i
		}
	}
	if currentMaxProfitIdx == -1 {
		return 0
	}
	currentProfit := Profits[currentMaxProfitIdx]
	// remove Capital[current]
	Profits = append(Profits[:currentMaxProfitIdx], Profits[currentMaxProfitIdx+1:]...)
	Captical = append(Captical[:currentMaxProfitIdx], Captical[currentMaxProfitIdx+1:]...)
	return loop(k-1, w+currentProfit, Profits, Captical) + currentProfit
}
