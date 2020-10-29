package solution

import "strconv"

var hour, minute int
var n_ int
var resu []string
var current []bool

func readBinaryWatch(n int) []string {
	hour = 4
	minute = 6
	n_ = n
	current = make([]bool, hour+minute)
	resu = []string{}
	dfs(0)
	return resu
}

func dfs(startIdx int) {
	if n_ < 0 {
		return
	}
	if startIdx == len(current) && n_ == 0 {
		if s, ok := translateBinTime(current); ok {
			resu = append(resu, s)
		}
		return
	}
	if startIdx >= len(current) {
		return
	}
	n_--
	current[startIdx] = true
	dfs(startIdx + 1)
	n_++
	current[startIdx] = false
	dfs(startIdx + 1)
}

func translateBinTime(ts []bool) (string, bool) {
	if len(ts) != len(current) {
		return "", false
	}
	currentHour := 0
	base2 := 1
	for i := hour - 1; i >= 0; i-- {
		if ts[i] {
			currentHour += base2
		}
		base2 <<= 1
	}
	if currentHour > 11 {
		return "", false
	}
	base2 = 1
	currentMinute := 0
	for i := len(current) - 1; i >= hour; i-- {
		if ts[i] {
			currentMinute += base2
		}
		base2 <<= 1
	}
	if currentMinute > 59 {
		return "", false
	}
	return strconv.Itoa(currentHour) + ":" + paddingMinute(strconv.Itoa(currentMinute)), true
}

func paddingMinute(s string) string {
	if len(s) == 2 {
		return s
	}
	return "0" + s
}
