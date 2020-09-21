package solution

var current []int
var resu [][]int
var n_, k_ int

func combine(n, k int) [][]int {
	resu = [][]int{}
	current = []int{}
	n_ = n
	k_ = k
	loop(0)
	return resu
}

func loop(startIdx int) {
	if len(current) == k_ {
		temp := make([]int, k_)
		copy(temp, current)
		resu = append(resu, temp)
		return
	}
	for i := startIdx; i < n_; i++ {
		current = append(current, i)
		loop(i + 1)
		current = current[:len(current)-1]
	}
}
