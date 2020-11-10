package solution

func lastStoneWeightII(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[1]
	}
	totalWeight := 0
	for i := 0; i < len(stones); i++ {
		totalWeight += stones[i]
	}
	// dp means min Weight left for j weight
	dp := make([]int, totalWeight>>1+1)
	for i := 0; i < totalWeight>>1+1; i++ {
		if i >= stones[0] {
			dp[i] = i - stones[0]
			continue
		}
		dp[i] = i
	}
	for i := 1; i < len(stones); i++ {
		for j := totalWeight >> 1; j > 0; j-- {
			if j >= stones[i] {
				dp[j] = min(dp[j], dp[j-stones[i]])
			}
		}
		// println(stones[i])
		// for k := 0; k < len(dp); k++ {
		// 	print(dp[k], " ")
		// }
		// println()
	}
	return totalWeight - (totalWeight>>1-dp[totalWeight>>1])<<1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
