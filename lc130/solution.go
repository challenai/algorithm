package solution

func solve(board [][]byte) {
	// idea: dsu
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	conn := make([][][]int, len(board))
	for i := 0; i < len(board); i++ {
		temp := make([][]int, len(board[i]))
		conn[i] = temp
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			conn[i][j] = []int{i, j}
		}
	}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] == 'O' {
			}
		}
	}
}

func solve2(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'P'
		}
		if board[i][len(board[0])-1] == 'O' {
			board[i][len(board[0])-1] = 'P'
		}
	}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			board[0][i] = 'P'
		}
		if board[len(board)-1][i] == 'O' {
			board[len(board)-1][i] = 'P'
		}
	}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' && (board[i-1][j] == 'P' || board[i][j-1] == 'P') {
				board[i][j] = 'P'
			}
		}
	}

	for i := len(board) - 2; i > 0; i-- {
		for j := len(board[i]) - 2; j > 0; j-- {
			if board[i][j] == 'O' && (board[i+1][j] == 'P' || board[i][j+1] == 'P') {
				board[i][j] = 'P'
			}
		}
	}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'P' {
				board[i][j] = 'O'
			}
		}
	}
}
