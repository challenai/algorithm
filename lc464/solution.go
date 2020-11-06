package solution

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	return maxChoosableInteger%desiredTotal == 0
}
