package solution

import "strconv"

func convertToBase7(num int) string {
	minus := num < 0
	if num < 0 {
		num = -num
	}
	rsu := ""
	for num > 0 {
		rsu = strconv.Itoa(num%7) + rsu
		num /= 7
	}
	if minus {
		rsu = "-" + rsu
	}
	return rsu
}
