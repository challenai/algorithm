package solution

var current, rsu int
var nums_ []int

func maxCoins(nums []int) int {
	current = 0
	rsu = 0
	nums_ = nums
	dfs()
	return rsu
}

func dfs() {
	if len(nums_) == 1 {
		if current+nums_[0] > rsu {
			rsu = current + nums_[0]
		}
		return
	}
	for i := 0; i < len(nums_); i++ {
		var point int
		if i == 0 {
			point = nums_[i] * nums_[i+1]
		} else if i == len(nums_)-1 {
			point = nums_[i] * nums_[i-1]
		} else {
			point = nums_[i] * nums_[i-1] * nums_[i+1]
		}
		current += point
		temp := make([]int, len(nums_))
		copy(temp, nums_)
		copy(nums_[i:], nums_[i+1:])
		nums_ = nums_[:len(nums_)-1]
		dfs()
		nums_ = temp
		current -= point
	}
}
