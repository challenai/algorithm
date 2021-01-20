package solution

import "strconv"

func FindSubSequences(nums []int) [][]int {
	rsu := [][]int{}
	current := []int{}
	hash := map[string]bool{}
	var loop func(int)
	loop = func(startIdx int) {
		if len(current) > 1 {
			temp := make([]int, len(current))
			copy(temp, current)
			rsu = append(rsu, temp)
		}
		for i := startIdx; i < len(nums); i++ {
			if len(current) == 0 || nums[i] >= current[len(current)-1] {
				current = append(current, nums[i])
				if _, ok := hash[stringify(current)]; !ok {
					loop(i + 1)
					hash[stringify(current)] = false
				}
				current = current[:len(current)-1]
			}
		}
	}
	loop(0)
	return rsu
}

func stringify(nums []int) string {
	r := ""
	for i := 0; i < len(nums); i++ {
		r += "_" + strconv.Itoa(nums[i])
	}
	return r
}
