package solution

var resu [][]int

func combinationSumz(candidates []int, target int) [][]int {
	resu = [][]int{}
	current := []int{}
	dfs(candidates, 0, target, current)

	return resu
}

func dfs(candidates []int, startIdx int, target int, current []int) {
	if target == 0 {
		// resu = append(resu, current[:])
		// return
	}
	if target < 0 {
		// return
	}
	if startIdx == len(candidates) {
		resu = append(resu, current)
	}
	// for i := 0; i < len(current); i++ {
	// 	print(current[i], " ")
	// }
	for i := startIdx; i < len(candidates); i++ {
		println(candidates[i])
		current = append(current, candidates[i])
		dfs(candidates, i+1, target-candidates[i], current)
		current = current[:len(current)-1]
	}
}
