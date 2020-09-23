package solution

func numDecodings(s string) int {
	var prev, temp, current, ptr int
	ptr = 1
	prev = 1
	current = 1
	for ptr < len(s) {
		temp = current
		if (s[ptr-1] == '1' && s[ptr] != '0') || (s[ptr-1] == '2' && s[ptr] <= '6' && s[ptr] >= '1') {
			current += prev
		}
		prev = temp
		ptr++
	}

	return current
}
