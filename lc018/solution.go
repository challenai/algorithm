package solution

import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	resu := [][]int{}
	resuHash := map[string][]int{}
	twoSums := map[int][][2]int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if idxs, ok := twoSums[target-nums[i]-nums[j]]; ok {
				for k := 0; k < len(idxs); k++ {
					entry := []int{idxs[k][0], idxs[k][1], i, j}
					sort.Slice(entry, func(a, b int) bool {
						return entry[a] < entry[b]
					})
					if checkNoDulplicate(entry) {
						// resu = append(resu, entry)
						key := strconv.Itoa(nums[entry[0]]) + "_" + strconv.Itoa(nums[entry[1]]) + "_" + strconv.Itoa(nums[entry[2]]) + "_" + strconv.Itoa(nums[entry[3]])
						if _, ok := resuHash[key]; !ok {
							resuHash[key] = []int{nums[idxs[k][0]], nums[idxs[k][1]], nums[i], nums[j]}
						}
					}
				}
			}
			if _, ok := twoSums[nums[i]+nums[j]]; ok {
				twoSums[nums[i]+nums[j]] = append(twoSums[nums[i]+nums[j]], [2]int{i, j})
			} else {
				twoSums[nums[i]+nums[j]] = [][2]int{
					{i, j},
				}
			}
		}
	}

	for _, idxs := range resuHash {
		resu = append(resu, idxs)
	}

	return resu
}

func checkNoDulplicate(nums []int) bool {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return false
		}
		hash[nums[i]] = false
	}
	return true
}
