package solution

import (
	"strconv"
)

var formula []string
var resu []int

func diffWaysToCompute(s string) []int {
	formula = []string{}
	resu = []int{}
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			current += string(s[i])
			continue
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' {
			if current == "" {
				return []int{}
			}
			formula = append(formula, current)
			current = ""
			formula = append(formula, string(s[i]))
		}
	}
	formula = append(formula, current)
	dfs()
	// for i := 0; i < len(resu); i++ {
	// 	print(resu[i], " ")
	// }
	// ln()
	return resu
}

func dfs() {
	if len(formula) == 1 {
		resu = append(resu, toInt(formula[0]))
	}
	for i := 1; i+1 < len(formula); i += 2 {
		temp := make([]string, len(formula))
		copy(temp, formula)
		// Todo: caculate formula
		formula[i-1] = getResult(formula[i-1], formula[i], formula[i+1])
		copy(formula[i:], formula[i+2:])
		formula = formula[:len(formula)-2]
		dfs()
		formula = temp
	}
}

func toInt(s string) int {
	resu, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return resu
}

func getResult(a, symbol, b string) string {
	a_ := toInt(a)
	b_ := toInt(b)
	if symbol == "+" {
		return strconv.Itoa(a_ + b_)
	}
	if symbol == "-" {
		return strconv.Itoa(a_ - b_)
	}
	return strconv.Itoa(a_ * b_)
}
