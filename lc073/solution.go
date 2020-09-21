package solution

func setZeroes(matrix [][]int) {
	var top, left bool
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			top = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			left = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
		if left {
			matrix[i][0] = 0
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
		if top {
			matrix[0][i] = 0
		}
	}
}
