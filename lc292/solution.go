package solution

func canWinNim(n int) bool {
	n %= 4
	return n != 0
}
