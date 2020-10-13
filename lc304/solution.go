package solution

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	temp := NumMatrix{
		sums: make([][]int, len(matrix)),
	}
	for i := 0; i < len(matrix); i++ {
		temp.sums[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 && j == 0 {
				temp.sums[i][j] = matrix[i][j]
			} else if i == 0 {
				temp.sums[i][j] = matrix[i][j] + temp.sums[i][j-1]
			} else if j == 0 {
				temp.sums[i][j] = matrix[i][j] + temp.sums[i-1][j]
			} else {
				temp.sums[i][j] = temp.sums[i-1][j] + temp.sums[i][j-1] + matrix[i][j] - temp.sums[i-1][j-1]
			}
		}
	}
	return temp
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sums[row2][col2]
	}
	if row1 == 0 {
		return this.sums[row2][col2] - this.sums[row2][col1-1]
	}
	if col1 == 0 {
		return this.sums[row2][col2] - this.sums[row1-1][col2]
	}
	return this.sums[row2][col2] + this.sums[row1-1][col1-1] - this.sums[row1-1][col2] - this.sums[row2][col1-1]
}
