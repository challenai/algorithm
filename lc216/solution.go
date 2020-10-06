package solution

var resu [][]int
var current []int
var currentSum, n_, k_ int

func combinationSum3(k, n int) [][]int {
	resu = [][]int{}
	current = []int{}
	currentSum = 0
	if n > 45 || n < 1 {
		return resu
	}
	n_ = n
	k_ = k
	dfs(1)
	return resu
}

func dfs(startIdx int) {
	if currentSum == n_ && len(current) == k_ {
		temp := make([]int, len(current))
		copy(temp, current)
		resu = append(resu, temp)
		return
	}
	if currentSum > n_ {
		return
	}
	for i := startIdx; i <= min(n_, 9); i++ {
		current = append(current, i)
		currentSum += i
		dfs(i + 1)
		current = current[:len(current)-1]
		currentSum -= i
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
