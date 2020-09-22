package solution

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	if len(newInterval) == 0 {
// 		return intervals
// 	}
// 	if len(intervals) == 0 {
// 		return []int{newInterval}
// 	}
// 	// bin search node[0] position
// 	var left, right, mid, temp int
// 	left = 0
// 	right = len(intervals) - 1
// 	for left < right-1 {
// 		mid = (left + right) / 2
// 		if intervals[mid][0] > newInterval {
// 			right = mid
// 		} else {
// 			left = mid
// 		}
// 	}
// 	temp = left
// 	if newInterval[1] <= intervals[left][1] {
// 		intervals[left][1] = newInterval[1]
// 	} else {
// 	}
// }

// merge every intervals in the process
func insert(intervals [][]int, newInterval []int) [][]int {
	// resu := [][]int{}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var left, right, mid int
	if intervals[0][0] > newInterval[0] {
		if intervals[0][1] > newInterval[1] {
			return intervals
		}
		intervals[0][0] = newInterval[0]
		intervals[0][1] = max(newInterval[1], intervals[0][1])
		left = 0
	} else if intervals[len(intervals)-1][1] < newInterval[1] {
		if intervals[len(intervals)-1][0] < newInterval[0] {
			return intervals
		}
		intervals[len(intervals)-1][1] = newInterval[1]
		intervals[len(intervals)-1][0] = min(newInterval[0], intervals[len(intervals)-1][0])
	} else {
		left = 0
		right = len(intervals) - 1
		for left < right-1 {
			mid = (left + right) / 2
			if newInterval[0] < intervals[mid][0] {
				right = mid
			} else {
				left = mid
			}
		}

		if intervals[left][1] >= newInterval[0] {
			intervals[left][1] = max(newInterval[1], intervals[left][1])
		} else if intervals[left+1][0] <= newInterval[1] {
			left++
			intervals[left][0] = min(newInterval[0], intervals[left][0])
		} else {
			return intervals
		}
	}
	i := left + 1
	for ; i < len(intervals)-1; i++ {
		if intervals[i][0] > intervals[left][1] {
			break
		}
		// bug: need compare and find a large one
		// intervals[left][1] = intervals[i][1]
		intervals[left][1] = max(intervals[i][1], intervals[left][1])
	}
	if i > left+1 {
		copy(intervals[left+1:], intervals[i:])
	}
	intervals = intervals[:len(intervals)-(i-left-1)]
	// for i := 0; i < len(intervals); i++ {
	// 	print(intervals[i][0], intervals[i][1], " | ")
	// }
	// println()

	return intervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
