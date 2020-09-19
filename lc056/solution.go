package solution

import "sort"

func merge(intervals [][]int) [][]int {
	var resu [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resu = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= resu[len(resu)-1][1] {
			resu[len(resu)-1][1] = max(intervals[i][i], resu[len(resu)-1][1])
		} else {
			resu = append(resu, intervals[i])
		}
	}

	return resu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
