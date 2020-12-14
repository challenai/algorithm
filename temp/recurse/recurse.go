package recurse

var current0 []int
var rsu0 [][]int
var nums0_ []int

func Combine(nums []int) [][]int {
	nums0_ = nums
	rsu0 = [][]int{}
	current0 = []int{}
	dfs(0)
	return rsu0
}

func dfs(startIdx int) {
	temp := make([]int, len(current0))
	copy(temp, current0)
	rsu0 = append(rsu0, temp)
	for i := startIdx; i < len(nums0_); i++ {
		current0 = append(current0, nums0_[i])
		dfs(i + 1)
		current0 = current0[:len(current0)-1]
	}
}

var current1 []int
var rsu1 [][]int
var nums1_ []int
var hash map[int]bool

func Arrange(nums []int) [][]int {
	current1 = []int{}
	nums1_ = nums
	rsu1 = [][]int{}
	hash = map[int]bool{}
	dfs1()
	return rsu1
}

func dfs1() {
	if len(hash) == len(nums1_) {
		temp := make([]int, len(current1))
		copy(temp, current1)
		rsu1 = append(rsu1, temp)
	}
	for i := 0; i < len(nums1_); i++ {
		if _, ok := hash[i]; !ok {
			current1 = append(current1, nums1_[i])
			hash[i] = false
			dfs1()
			current1 = current1[:len(current1)-1]
			delete(hash, i)
		}
	}
}
