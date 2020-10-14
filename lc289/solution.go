package solution

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			temp := cntLives(board, i, j)
			if i == 1 && j == 0 {
			}
			if board[i][j] == 0 && temp == 3 {
				board[i][j] = -1
				continue
			}
			if board[i][j] == 1 && (temp < 2 || temp > 3) {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}

func cntLives(board [][]int, i, j int) int {
	rsu := 0
	rsu += getPosInBoard(board, i-1, j-1)
	rsu += getPosInBoard(board, i-1, j+1)
	rsu += getPosInBoard(board, i-1, j)
	rsu += getPosInBoard(board, i, j-1)
	rsu += getPosInBoard(board, i, j+1)
	rsu += getPosInBoard(board, i+1, j-1)
	rsu += getPosInBoard(board, i+1, j+1)
	rsu += getPosInBoard(board, i+1, j)
	return rsu
}

func getPosInBoard(board [][]int, i, j int) int {
	if i >= 0 && i < len(board) && j < len(board[i]) && j >= 0 {
		if board[i][j] == 1 || board[i][j] == 2 {
			return 1
		}
	}
	return 0
}
