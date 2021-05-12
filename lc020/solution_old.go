package solution

func isValid(s string) bool {
	// idea: stack couple compare
	st := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		} else if s[i] == ']' {
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		} else if s[i] == '}' {
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}
	return len(st) == 0
}
