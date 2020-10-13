package solution

import (
	"strconv"
)

func calculate(s string) int {
	var resu int
	var err error

	q := []string{}
	current := ""
	// clean the formulation
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			current += string(s[i])
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
	// for i := 0; i < len(q); i++ {
	// 	print(q[i])
	// }
	// println()
	resu, err = strconv.Atoi(q[0])
	if err != nil || len(q) == 0 {
		return 0
	}
	q2 := []int{}
	for i := 1; i+1 < len(q); i += 2 {
		if q[i] == "+" {
			q2 = append(q2, resu)
			resu, err = strconv.Atoi(q[i+1])
			if err != nil {
				return 0
			}
		}
		if q[i] == "-" {
			q2 = append(q2, resu)
			resu, err = strconv.Atoi(q[i+1])
			if err != nil {
				return 0
			}
			resu = -resu
		}
		if q[i] == "*" {
			temp, err := strconv.Atoi(q[i+1])
			if err != nil {
				return 0
			}
			resu *= temp
			continue
		}
		if q[i] == "/" {
			temp, err := strconv.Atoi(q[i+1])
			if err != nil {
				return 0
			}
			resu /= temp
			continue
		}
	}
	if resu != 0 {
		q2 = append(q2, resu)
	}

	resu = 0
	for i := 0; i < len(q2); i++ {
		// print(q2[i], " ")
		resu += q2[i]
	}

	// println(resu)
	return resu
}
