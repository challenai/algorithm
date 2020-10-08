package solution

func maximalSquare(matrix [][]int) int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// if i == 0 && j == 0 {
			// 	continue
			// } else if i == 0 {
			// 	matrix[i][j] += matrix[i][j-1]
			// } else if j == 0 {
			// 	matrix[i][j] += matrix[i-1][j]
			// } else {
			// 	matrix[i][j] += matrix[i-1][j] + matrix[i][j-1]
			// }
		}
	}
	// current len of square
	current := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// if current == 0 && matrix[i][j] != 0 {
			// 	current = 1
			// }
			println(matrix[i][j], matrix[i-current][j], matrix[i][j-current])
			for i+1 > current && j+1 > current && matrix[i][j]+matrix[i-current][j-current]-matrix[i-current][j]-matrix[i][j-current] >= (current+1)*(current+1) {
				current++
			}
		}
	}
	println(current)

	return current * current
}
