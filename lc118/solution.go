package solution

func generate(numRows int) [][]int {
	resu := [][]int{{1}}
	if numRows == 1 {
		return resu
	}
	resu = append(resu, []int{1, 1})
	if numRows == 2 {
		return resu
	}
	counter := 3
	for counter <= numRows {
		resu = append(resu, []int{1})
		for i := 1; i < len(resu[len(resu)-2]); i++ {
			resu[len(resu)-1] = append(resu[len(resu)-1], resu[len(resu)-2][i]+resu[len(resu)-2][i-1])
		}
		resu[len(resu)-1] = append(resu[len(resu)-1], 1)
		counter++
	}

	return resu
}
