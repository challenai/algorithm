package solution

func countBattleships(board [][]byte) int {
	rsu := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				rsu++
				paddingConnection(board, i, j)
			}
		}
	}
	return rsu
}

func paddingConnection(board [][]byte, m, n int) {
	for i := m; i < len(board); i++ {
		if board[i][n] == '.' {
			break
		}
		board[i][n] = '.'
	}
	for i := n + 1; i < len(board[0]); i++ {
		if board[m][i] == '.' {
			break
		}
		board[m][i] = '.'
	}
}
