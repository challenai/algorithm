package solution

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return []int{newInterval}
	}
	// bin search node[0] position
	var left, right, mid, temp int
	left = 0
	right = len(intervals) - 1
	for left < right-1 {
		mid = (left + right) / 2
		if intervals[mid][0] > newInterval {
			right = mid
		} else {
			left = mid
		}
	}
	temp = left
	if newInterval[1] <= intervals[left][1] {
		intervals[left][1] = newInterval[1]
	} else {
	}
}
