package solution

import (
	"sort"
	"strconv"
)

var resu [][]int
var current []int
var resuHash map[string]bool
var nums_ []int

func subsetsWithDup(nums []int) [][]int {
	resu = [][]int{}
	current = []int{}
	nums_ = nums
	sort.Slice(nums_, func(i, j int) bool {
		return nums_[i] < nums_[j]
	})
	resuHash = map[string]bool{}
	loop(0)

	return resu
}

func loop(startIdx int) {
	key := serialize(current)
	if _, ok := resuHash[key]; !ok {
		resuHash[key] = false
		temp := make([]int, len(current))
		copy(temp, current)
		resu = append(resu, temp)
	}
	for i := startIdx; i < len(nums_); i++ {
		current = append(current, nums_[i])
		loop(i + 1)
		current = current[:len(current)-1]
	}
}

func serialize(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	resu := strconv.Itoa(nums[0])
	for i := 1; i < len(nums); i++ {
		resu += "_" + strconv.Itoa(nums[i])
	}
	return resu
}
