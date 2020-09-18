package solution

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}
	idx := 0
	for idx < len(matrix)-1-idx {
		for i := idx; i <= len(matrix)-2-idx; i++ {
			matrix[idx][i], matrix[i][len(matrix)-1-idx] = matrix[i][len(matrix)-1-idx], matrix[idx][i]
			matrix[idx][i], matrix[len(matrix)-1-idx][len(matrix)-1-i] = matrix[len(matrix)-1-idx][len(matrix)-1-i], matrix[idx][i]
			matrix[idx][i], matrix[len(matrix)-1-i][idx] = matrix[len(matrix)-1-i][idx], matrix[idx][i]
		}
		idx++
	}
}
