package solution

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	lenRecorder := len(tokens)
	for len(tokens) > 1 {
		for i := 0; i < len(tokens); i++ {
			if isSymbol(tokens[i]) {
				if i > 1 && !isSymbol(tokens[i-1]) && !isSymbol(tokens[i-2]) {
					temp2 := parseInt(tokens[i-1])
					temp1 := parseInt(tokens[i-2])
					if tokens[i] == "+" {
						tokens[i-2] = strconv.Itoa(temp1 + temp2)
					} else if tokens[i] == "-" {
						tokens[i-2] = strconv.Itoa(temp1 - temp2)
					} else if tokens[i] == "*" {
						tokens[i-2] = strconv.Itoa(temp1 * temp2)
					} else {
						tokens[i-2] = strconv.Itoa(temp1 / temp2)
					}
					copy(tokens[i-1:], tokens[i+1:])
					tokens = tokens[:len(tokens)-2]
					break
				}
			}
		}
		if lenRecorder == len(tokens) {
			break
		}
	}
	return parseInt(tokens[0])
}

func isSymbol(s string) bool {
	if s == "*" || s == "+" || s == "-" || s == "/" {
		return true
	}
	return false
}

func parseInt(s string) int {
	rsu, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return rsu
}
