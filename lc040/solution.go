package solution

import "sort"

var resu [][]int
var current []int

func combinationSumz(candidates []int, target int) [][]int {
	resu = [][]int{}
	current = []int{}
	sort.Slice(candidates, func(a, b int) bool {
		return candidates[a] < candidates[b]
	})
	dfs(candidates, 0, target)

	return resu
}

func dfs(candidates []int, startIdx int, target int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		resu = append(resu, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := startIdx; i < len(candidates); i++ {
		// awesome condition, i > startIdx..., I used i > 0 before. but stopped in the next similar element
		if i > startIdx && candidates[i] == candidates[i-1] {
			continue
		}
		current = append(current, candidates[i])
		dfs(candidates, i+1, target-candidates[i])
		current = current[:len(current)-1]
	}
}
