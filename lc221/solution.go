package solution

// func maximalSquare(matrix [][]int) int {
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			// if i == 0 && j == 0 {
// 			// 	continue
// 			// } else if i == 0 {
// 			// 	matrix[i][j] += matrix[i][j-1]
// 			// } else if j == 0 {
// 			// 	matrix[i][j] += matrix[i-1][j]
// 			// } else {
// 			// 	matrix[i][j] += matrix[i-1][j] + matrix[i][j-1]
// 			// }
// 		}
// 	}
// 	// current len of square
// 	current := 0
// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[i]); j++ {
// 			// if current == 0 && matrix[i][j] != 0 {
// 			// 	current = 1
// 			// }
// 			println(matrix[i][j], matrix[i-current][j], matrix[i][j-current])
// 			for i+1 > current && j+1 > current && matrix[i][j]+matrix[i-current][j-current]-matrix[i-current][j]-matrix[i][j-current] >= (current+1)*(current+1) {
// 				current++
// 			}
// 		}
// 	}
// 	println(current)

// 	return current * current
// }

func maximalSquare(matrix [][]int) int {
	// use template descent stack ...
	// forgive me I can just use this method...
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	var st []int
	resu := 0
	for i := 0; i < len(matrix); i++ {
		st = []int{-1}
		for j := 0; j < len(matrix[i]); j++ {
			for len(st) > 1 && matrix[i][j] <= st[len(st)-1] {
				temp := st[len(st)-1]
				st = st[:len(st)-1]
				if newMax := sqrt2(min(j-st[len(st)-1]-1, matrix[i][temp])); newMax > resu {
					resu = newMax
				}
			}
			st = append(st, j)
		}
		for len(st) > 1 {
			temp := st[len(st)-1]
			st = st[:len(st)-1]
			if newMax := sqrt2(min(temp-st[len(st)-1], matrix[i][temp])); newMax > resu {
				resu = newMax
			}
		}
	}
	return resu
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sqrt2(a int) int {
	return a * a
}
