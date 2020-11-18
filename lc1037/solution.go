package solution

func isBoomerang(points [][]int) bool {
	if len(points) != 3 {
		return false
	}
	if points[0][0] == points[1][0] && points[2][0] == points[1][0] {
		return false
	}
	if points[0][0] == points[1][0] || points[2][0] == points[1][0] || points[0][0] == points[2][0] {
		return true
	}

	dv1, rest1, _ := diffPoints(points[0], points[1])
	dv2, rest2, _ := diffPoints(points[1], points[2])
	return dv1 != dv2 || rest1 != rest2
}

func diffPoints(point1, point2 []int) (dv, rest, symb int) {
	temp1 := point1[1] - point2[1]
	temp2 := point1[0] - point2[0]
	if temp1 < 0 && temp2 < 0 {
		temp1 = -temp1
		temp2 = -temp2
		symb = 1
	} else if temp1 < 0 {
		temp1 = -temp1
		symb = -1
	} else if temp2 < 0 {
		temp2 = -temp2
		symb = -1
	} else {
		symb = 1
	}

	dv = temp1 / temp2
	rest = temp1 % temp2
	return
}
