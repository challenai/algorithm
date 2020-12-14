package recurse

var resu3 [][]int
var nums3_, current3 []int
var hash3 map[int]bool

func Arrange3(nums []int) [][]int {
	nums3_ = nums
	current3 = []int{}
	resu3 = [][]int{}
	hash3 = map[int]bool{}
	dfs3()
	return resu3
}

func dfs3() {
	if len(hash) == len(nums3_) {
		temp := make([]int, len(current3))
		copy(temp, current3)
		resu3 = append(resu3, temp)
	}
	for i := 0; i < len(nums3_); i++ {
		if _, ok := hash3[i]; !ok {
			current3 = append(current3, nums3_[i])
			hash3[i] = false
			dfs3()
			current3 = current3[:len(current3)-1]
			delete(hash3, i)
		}
	}
}

var nums4_ []int
var resu4 [][]int
var current4 []int

func Combine3(nums []int) [][]int {
	resu4 = [][]int{}
	nums4_ = nums
	current4 = []int{}
	dfs4(0)
	return resu4
}

func dfs4(startIdx int) {
	temp := make([]int, len(current4))
	copy(temp, current4)
	resu4 = append(resu4, temp)
	for i := startIdx; i < len(nums4_); i++ {
		current4 = append(current4, nums4_[i])
		dfs4(i + 1)
		current4 = current4[:len(current4)-1]
	}
}
