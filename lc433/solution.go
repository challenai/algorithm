package solution

var potiential map[string]bool
var rsu, bankLen int
var end_ string

func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	potiential = map[string]bool{}
	for i := 0; i < len(bank); i++ {
		if bank[i] != start {
			potiential[bank[i]] = false
		}
	}
	bankLen = len(bank)
	rsu = 0
	end_ = end
	dfs(start)
	return rsu
}

func dfs(start string) {
	if start == end_ {
		if rsu == 0 || rsu > bankLen-len(potiential) {
			rsu = bankLen - len(potiential)
		}
		return
	}
	if rsu != 0 && bankLen-len(potiential) > rsu {
		return
	}
	for k := range potiential {
		if checkTransmitable(start, k) {
			delete(potiential, k)
			dfs(k)
			potiential[k] = true
		}
	}
}

func checkTransmitable(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	unsame := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			unsame++
		}
	}
	return unsame == 1
}
