package solution

var rsu, m_ int
var nums_, groupSum []int

func splitArray(nums []int, m int) int {
	rsu = 0
	m_ = m
	nums_ = nums
	groupSum = make([]int, m)
	dfs(0, 0)
	println(rsu)
	return rsu
}

func dfs(idx, groupId int) {
	if idx == len(nums_) && groupId == m_ {
		if temp := max(groupSum); rsu == 0 || rsu > temp {
			rsu = temp
		}
		return
	}
	if groupId >= m_ || idx >= len(nums_) {
		return
	}
	groupSum[groupId] += nums_[idx]
	dfs(idx+1, groupId)
	dfs(idx+1, groupId+1)
	groupSum[groupId] -= nums_[idx]
}

func max(list []int) int {
	if len(list) == 0 {
		return 0
	}
	resu := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > resu {
			resu = list[i]
		}
	}
	return resu
}
