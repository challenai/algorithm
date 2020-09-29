package solution

var resu [][]int
var current []int
var nums_ []int

// func subsets(nums []int) [][]int {
// 	resu = [][]int{}
// 	current = []int{}
// 	nums_ = nums
// 	loop(0)
// 	return resu
// }

// func loop(startIdx int) {
// 	temp := make([]int, len(current))
// 	copy(temp, current)
// 	resu = append(resu, temp)
// 	for i := startIdx; i < len(nums_); i++ {
// 		current = append(current, nums_[i])
// 		loop(i + 1)
// 		current = current[:len(current)-1]
// 	}
// }

func subsets2(nums []int) [][]int {
	resu = [][]int{}
	nums_ = nums
	current = []int{}
	loop2(0)
	return resu
}

func loop2(startIdx int) {
	temp := make([]int, len(current))
	copy(temp, current)
	resu = append(resu, temp)
	for i := startIdx; i < len(nums_); i++ {
		current = append(current, nums_[i])
		loop2(i + 1)
		current = current[:len(current)-1]
	}
}
