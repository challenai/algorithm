package solution

import "strconv"

func calculate(s string) int {
	q := []string{}
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			current += string(s[i])
			continue
		}
		if s[i] == ' ' {
			continue
		}
		if current != "" {
			q = append(q, current)
			current = ""
		}
		q = append(q, string(s[i]))
	}
	if current != "" {
		q = append(q, current)
	}
	st := []int{}
	sum := 0
	for i := 0; i < len(q); i++ {
		if q[i] == "(" {
			if i == 0 || q[i-1] == "+" {
				st = append(st, 1)
				st = append(st, sum)
			} else {
				st = append(st, -1)
				st = append(st, sum)
			}
			sum = 0
			continue
		}
		if q[i] == ")" {

		}
	}
	for i := 0; i < len(q); i++ {
		print(q[i])
	}
	println()
	return 0
}

func getNum(numStr, symbol string) int {
	resu, _ := strconv.Atoi(numStr)
	if symbol == "-" {
		return -resu
	}
	return resu
}
