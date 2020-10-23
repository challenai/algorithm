package solution

func maxSumSubmatrix(matrix [][]int, k int) int {
	// idea: translate the original matrix to a new sum matrix
	// traversal all the possible matrix combination
	// compare the result
	// tc is O(mn^2)
	resu := k + 1
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i-1][j]
			if resu == k+1 && matrix[i][j] <= k {
				resu = matrix[i][j]
			}
		}
	}
	if resu == k {
		return k
	}
	if resu == k+1 {
		return 0
	}
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	var current int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			for z := 0; z < len(matrix); z++ {
				for l := 0; l < len(matrix[z]); l++ {
					current = caculateSum(matrix, i, j, z, l)
					if current <= k && current > resu {
						resu = current
					}
				}
			}
		}
	}
	return resu
}

func caculateSum(matrix [][]int, i, j, z, l int) int {
	if i == 0 && j == 0 {
		return matrix[z][l]
	}
	if i == 0 {
		return matrix[z][l] - matrix[z][j]
	}
	if j == 0 {
		return matrix[z][l] - matrix[i][l]
	}
	return matrix[z][l] + matrix[z][l] - matrix[z][j] - matrix[i][l]
}
