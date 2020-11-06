package solution

var options [][]int
var currentSum, rsu int

func fourSumCount(a, b, c, d []int) int {
	options = [][]int{a, b, c, d}
	rsu = 0
	currentSum = 0
	dfs(0)
	return rsu
}

func dfs(idx int) {
	if idx == len(options) {
		if currentSum == 0 {
			rsu++
		}
		return
	}
	for i := 0; i < len(options[idx]); i++ {
		currentSum += options[idx][i]
		dfs(idx + 1)
		currentSum -= options[idx][i]
	}
}
