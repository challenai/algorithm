package solution

import (
	"fmt"
	"strconv"
)

var resu [][]int
var dp [][]int
var existPath map[string]bool
var matrix_ [][]int

func pacificAtlantic(matrix [][]int) [][]int {
	matrix_ = matrix
	resu = [][]int{}
	dp = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	existPath = map[string]bool{}
	for i := 0; i < len(matrix); i++ {
		dfs(0, i, 0)
	}
	for i := 0; i < len(matrix[0]); i++ {
		dfs(i, 0, 0)
	}
	existPath = map[string]bool{}
	for i := 0; i < len(matrix); i++ {
		dfs(len(matrix)-1, i, 1)
	}
	for i := 0; i < len(matrix[0]); i++ {
		dfs(i, len(matrix[0])-1, 1)
	}
	// for i := 0; i < len(resu); i++ {
	// 	print(resu[i][0], " ", resu[i][1], " | ")
	// }
	// println()
	// for i := 0; i < len(dp); i++ {
	// 	for j := 0; j < len(dp[i]); j++ {
	// 		print(dp[i][j], " ")
	// 	}
	// 	println()
	// }
	// println()
	return resu
}

func dfs(i, j, ocean int) {
	key := fmt.Sprintf("%s_%s", strconv.Itoa(i), strconv.Itoa(j))
	if _, ok := existPath[key]; ok {
		return
	}
	existPath[key] = false
	dp[i][j]++
	if dp[i][j] == 2 {
		resu = append(resu, []int{i, j})
	}

	if j > 0 && matrix_[i][j] <= matrix_[i][j-1] {
		dfs(i, j-1, ocean)
	}
	if i > 0 && matrix_[i][j] <= matrix_[i-1][j] {
		dfs(i-1, j, ocean)
	}
	if i < len(matrix_)-1 && matrix_[i][j] <= matrix_[i+1][j] {
		dfs(i+1, j, ocean)
	}
	if j < len(matrix_[0])-1 && matrix_[i][j] <= matrix_[i][j+1] {
		dfs(i, j+1, ocean)
	}
}
