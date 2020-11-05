package solution

import "sort"

func findMinArrowShots(points [][]int) int {
	rsu := 0
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	bursted := map[int]bool{}
	// 反正每个球总是要被扎的，要扎就尽尽可能多的。。。
	for i := 0; i < len(points); i++ {
		if _, ok := bursted[i]; ok {
			continue
		}
		for j := i + 1; j < len(points); j++ {
			if points[i][1] >= points[j][0] {
				bursted[j] = true
			} else {
				break
			}
		}
		rsu++
	}
	return rsu
}
