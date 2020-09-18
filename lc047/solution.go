package solution

import (
	"sort"
	"strconv"
)

var resu [][]int
var current, nums_ []int
var numsSet map[int]bool
var resultsSet map[string]bool

func permuteUnique(nums []int) [][]int {
	resu = [][]int{}
	current = []int{}
	nums_ = nums
	sort.Slice(nums_, func(i, j int) bool {
		return nums_[i] < nums_[j]
	})
	numsSet = map[int]bool{}
	resultsSet = map[string]bool{}
	loop()
	return resu
}

func loop() {
	if len(current) == len(nums_) {
		key := joinIntKeys(current, "_")
		if _, ok := resultsSet[key]; !ok {
			temp := make([]int, len(current))
			copy(temp, current)
			resu = append(resu, temp)
			resultsSet[key] = false
		}
	}
	for i := 0; i < len(nums_); i++ {
		if _, ok := numsSet[i]; !ok {
			current = append(current, nums_[i])
			numsSet[i] = false
			loop()
			current = current[:len(current)-1]
			delete(numsSet, i)
		}
	}
}

func joinIntKeys(current []int, seperator string) string {
	if len(current) == 0 {
		return ""
	}
	resu := strconv.Itoa(current[0])
	for i := 0; i < len(current); i++ {
		resu += "_" + strconv.Itoa(current[i])
	}
	return resu
}
