package solution

import "sort"

func findContentChildren(g, s []int) int {
	rsu := 0
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	gPtr := 0
	for i := 0; i < len(s); i++ {
		if gPtr >= len(g) {
			break
		}
		if s[i] >= g[gPtr] {
			rsu++
			gPtr++
		}
	}
	return rsu
}
