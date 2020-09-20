package solution

import "strconv"

func addBinary(a, b string) string {
	var ptr, needPlusOne, current int
	ptr = 0
	needPlusOne = 0
	resu := ""

	for ptr < len(a) || ptr < len(b) {
		current = 0
		if ptr < len(a) && ptr < len(b) {
			if a[len(a)-1-ptr] != b[len(b)-1-ptr] {
				current = 1
			} else if a[len(a)-1-ptr] == '1' {
				current = 2
			}
		} else if ptr < len(a) && a[len(a)-1-ptr] == '1' {
			current = 1
		} else if ptr < len(b) && b[len(b)-1-ptr] == '1' {
			current = 1
		}
		current += needPlusOne
		if current >= 2 {
			needPlusOne = 1
			current %= 2
		} else {
			needPlusOne = 0
		}
		resu = strconv.Itoa(current) + resu
		ptr++
	}
	if needPlusOne > 0 {
		return "1" + resu
	}
	return resu
}
