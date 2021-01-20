package solution

var current string
var hash map[byte]int
var step, rsu int

func FindMinStep(board, hand string) int {
	// dp := make([][]int, len(board))
	// for i := 0; i < len(board); i++ {
	// 	dp[i] = make([]int, len(board))
	// }
	rsu = -1
	step = 0
	hash = map[byte]int{}
	current = board
	for i := 0; i < len(hand); i++ {
		if _, ok := hash[hand[i]]; ok {
			hash[hand[i]]++
			continue
		}
		hash[hand[i]] = 1
	}
	loop()
	return rsu
}

func loop() {
	for i := 1; i < len(current); i++ {

	}
}
