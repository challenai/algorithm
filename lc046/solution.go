package solution

var resu [][]int
var nums_, current []int
var currentHash map[int]bool

func permute(nums []int) [][]int {
	resu = [][]int{}
	current = []int{}
	nums_ = nums
	currentHash = map[int]bool{}
	loop()
	return resu
}

func loop() {
	if len(current) == len(nums_) {
		temp := make([]int, len(current))
		copy(temp, current)
		resu = append(resu, temp)
		return
	}
	for i := 0; i < len(nums_); i++ {
		if _, ok := currentHash[nums_[i]]; !ok {
			current = append(current, nums_[i])
			currentHash[nums_[i]] = false
			loop()
			current = current[:len(current)-1]
			delete(currentHash, nums_[i])
		}
	}
}
