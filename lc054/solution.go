package solution

func spiralOrder(matrix [][]int) []int {
	resu := []int{}
	idx := 0
	for idx <= len(matrix)-1-idx && idx <= len(matrix[0])-1-idx {
		if idx <= len(matrix)-1-idx {
			for i := idx; i < len(matrix[0])-1-idx; i++ {
				resu = append(resu, matrix[idx][i])
			}
		}

		if idx <= len(matrix[0])-1-idx {
			for i := idx; i < len(matrix)-1-idx; i++ {
				resu = append(resu, matrix[i][len(matrix[0])-1-idx])
			}
		}

		if idx < len(matrix)-1-idx {
			for i := idx; i < len(matrix[0])-1-idx; i++ {
				resu = append(resu, matrix[len(matrix)-1-idx][len(matrix[0])-1-i])
			}
		}

		if idx < len(matrix[0])-1-idx {
			for i := idx; i < len(matrix)-1-idx; i++ {
				resu = append(resu, matrix[len(matrix)-1-i][idx])
			}
		}
		idx++
	}

	for i := 0; i < len(resu); i++ {
		print(resu[i], " ")
	}
	println()

	return resu
}
