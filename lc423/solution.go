package solution

import "strconv"

var byteCnts map[byte]int
var intToText []string
var rsu, current string

func originalDigits(s string) string {
	byteCnts = map[byte]int{}
	intToText = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	for i := 0; i < len(s); i++ {
		if _, ok := byteCnts[s[i]]; ok {
			byteCnts[s[i]]++
		} else {
			byteCnts[s[i]] = 1
		}
	}
	current = ""
	rsu = ""
	dfs(0)
	return rsu
}

func dfs(num int) {
	if len(byteCnts) == 0 {
		rsu = current
	}
	if num > 9 {
		return
	}
	if len(rsu) > 0 {
		return
	}
	availableLen := removeDigits(num)
	if availableLen == len(intToText[num]) {
		current += strconv.Itoa(num)
		dfs(num)
		current = current[:len(current)-1]
		addDigits(intToText[num], availableLen)
		return
	}
	addDigits(intToText[num], availableLen)
	dfs(num + 1)
}

func addDigits(s string, end int) {
	for i := 0; i < end; i++ {
		if _, ok := byteCnts[s[i]]; ok {
			byteCnts[s[i]]++
		} else {
			byteCnts[s[i]] = 1
		}
	}
}

func removeDigits(num int) int {
	numStr := intToText[num]
	for i := 0; i < len(numStr); i++ {
		if _, ok := byteCnts[numStr[i]]; !ok {
			return i
		} else {
			byteCnts[numStr[i]]--
			if byteCnts[numStr[i]] == 0 {
				delete(byteCnts, numStr[i])
			}
		}
	}
	return len(numStr)
}
