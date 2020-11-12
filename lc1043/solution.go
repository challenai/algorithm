package solution

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, k)
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		for zz := 0; zz < len(dp); zz++ {
			print(dp[zz], " ")
		}
		println()
		for j := 1; j < k; j++ {
			if i-j < 0 {
				continue
			}
			dp[j] = dp[j-1]
		}
		for j := 0; j < k; j++ {
			if i-j < 0 {
				continue
			}
			println("->", i, j)
			dp[0] = max(dp[0], dp[j]+checkMax(arr, i-j, i)*(j))
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// var nums []int
// var currentSum, k_, rsu, currentK, temp, tempMax int

// func maxSumAfterPartitioning(arr []int, k int) int {
// 	nums = arr
// 	k_ = k
// 	rsu = 0
// 	currentK = 0
// 	currentSum = 0
// 	dfs(0)
// 	return rsu
// }

// func dfs(idx int) {
// 	if idx == len(nums) {
// 		tempMax := checkMax(idx-currentK+1, idx-1) * (currentK - 1)
// 		currentSum += tempMax
// 		if rsu == 0 || currentSum > rsu {
// 			rsu = currentSum
// 		}
// 		currentSum -= tempMax
// 		return
// 	}
// 	if currentK < k_ {
// 		currentK++
// 		dfs(idx + 1)
// 		currentK--
// 	}

// 	temp := currentK
// 	tempMax := checkMax(idx-currentK+1, idx) * currentK
// 	currentSum += tempMax
// 	currentK = 1
// 	dfs(idx + 1)
// 	currentSum -= tempMax
// 	currentK = temp
// }

func checkMax(nums []int, start, end int) int {
	maxV := nums[end]
	for i := start; i < end; i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	return maxV
}
