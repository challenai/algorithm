package solution

var rsu, current int
var target_ int
var nums_ []int

func combinationSum4(nums []int, target int) int {
	nums_ = nums
	target_ = target
	rsu = 0
	dfs(0)
	return rsu
}

func dfs(startIdx int) {
	if current == target_ {
		rsu++
		return
	}
	if current > target_ {
		return
	}
	for i := startIdx; i < len(nums_); i++ {
		current += nums_[i]
		dfs(i)
		current -= nums_[i]
	}
}
