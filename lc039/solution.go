package solution

var resu [][]int

func combinationSum(candidates []int, target int) [][]int {
	// idea: use dp to handle bag problem. just a simple practice...
	resu = [][]int{}
	current := []int{}
	dfs(candidates, 0, target, current)
	return resu
}

func dfs(candidates []int, startIdx, target int, current []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		resu = append(resu, current[:])
		return
	}
	for i := startIdx; i < len(candidates); i++ {
		current = append(current, candidates[i])
		dfs(candidates, i, target-candidates[i], current)
		current = current[:len(current)-1]
	}
}
