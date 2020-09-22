package solution

func maximalRectangle(matrix [][]byte) int {
	heightsMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		heightsMatrix[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			heightsMatrix[0][i] = 1
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heightsMatrix[i][j] = heightsMatrix[i-1][j] + 1
			}
		}
	}
	// for i := 0; i < len(heightsMatrix); i++ {
	// 	for j := 0; j < len(heightsMatrix[0]); j++ {
	// 		print(heightsMatrix[i][j], " ")
	// 	}
	// 	println()
	// }
	// println()

	var temp int
	var st []int
	resu := 0
	for i := 0; i < len(heightsMatrix); i++ {
		st = []int{-1}
		for j := 0; j < len(heightsMatrix[i]); j++ {
			if len(st) > 1 && heightsMatrix[i][j] <= heightsMatrix[i][st[len(st)-1]] {
				for len(st) > 1 && heightsMatrix[i][j] <= heightsMatrix[i][st[len(st)-1]] {
					temp = st[len(st)-1]
					st = st[:len(st)-1]
					resu = max(resu, (j-st[len(st)-1]-1)*heightsMatrix[i][temp])
				}
			}
			st = append(st, j)
		}
		for len(st) > 1 {
			temp = st[len(st)-1]
			st = st[:len(st)-1]
			resu = max(resu, (temp-st[len(st)-1])*heightsMatrix[i][temp])
		}
	}

	return resu
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
