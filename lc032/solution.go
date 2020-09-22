package solution

func longestValidParentheses(str string) int {
	var counter, resu, current int
	resu = 0
	counter = 0
	current = 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			counter++
		} else if str[i] == ')' {
			if counter <= 0 {
				counter = 0
				current = 0
				continue
			}
			current++
			if current > resu {
				resu = current
			}
			counter--
		}
		// ignore no () chars...
	}
	return resu * 2
}
