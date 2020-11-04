package solution

import "sort"

var intervals [][]int
var rsu, current int
var currentErase []bool

func eraseOverlapIntervals(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i][0] == nums[j][0] {
			return nums[i][1] < nums[j][1]
		}
		return nums[i][0] < nums[j][0]
	})
	intervals = nums
	currentErase = make([]bool, len(nums))
	rsu = -1
	current = 0
	dfs(0)
	return rsu
}

func dfs(idx int) {
	if idx == len(intervals) {
		if rsu == -1 || rsu > current {
			rsu = current
		}
		return
	}
	// remove me
	if currentErase[idx] {
		dfs(idx + 1)
		return
	}
	currentErase[idx] = true
	current++
	dfs(idx + 1)
	current--
	currentErase[idx] = false
	temp := make([]bool, len(intervals))
	copy(temp, currentErase)
	temp2 := 0
	for i := idx + 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[idx][1] {
			break
		}
		// has overlap
		if !currentErase[i] {
			temp2++
			currentErase[i] = true
		}
	}
	current += temp2
	dfs(idx + 1)
	currentErase = temp
	current -= temp2
}
