package solution

import "sort"

func findRightInterval(intervals [][]int) []int {
	rsu := make([]int, len(intervals))
	mapper := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		mapper[i] = i
	}
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(mapper, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var low, high, mid int
	for i := 0; i < len(intervals); i++ {
		low = 0
		high = len(sortedIntervals) - 1
		if intervals[i][1] > sortedIntervals[high][0] {
			rsu[i] = -1
			continue
		}
		for low < high-1 {
			mid = low + (high-low)>>1
			if intervals[i][1] == sortedIntervals[mid][0] {
				high = mid
				low = high
				break
			} else if intervals[i][1] > intervals[mid][0] {
				low = mid
			} else {
				high = mid
			}
		}
		if mapper[high] != i {
			rsu[i] = mapper[high]
		}
	}
	return rsu
}
