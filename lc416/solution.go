package solution

var sum0, sum1 int
var nums_ []int
var rsu bool

func canPartition(nums []int) bool {
	nums_ = nums
	rsu = false
	sum0 = 0
	sum1 = 0
	dfs(0)
	return rsu
}

func dfs(idx int) {
	if rsu {
		return
	}
	if idx == len(nums_) && sum0 == sum1 {
		rsu = true
		return
	}

	if idx >= len(nums_) {
		return
	}
	sum0 += nums_[idx]
	dfs(idx + 1)
	sum0 -= nums_[idx]
	sum1 += nums_[idx]
	dfs(idx + 1)
	sum1 -= nums_[idx]
}
