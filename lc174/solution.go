package solution

// func calculateMinimumHP(dugeon [][]int) int {
// 	for i := 0; i < len(dugeon); i++ {
// 		for j := 0; j < len(dugeon[i]); j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			if i == 0 {
// 				dugeon[i][j] += dugeon[i][j-1]
// 			} else if j == 0 {
// 				dugeon[i][j] += dugeon[i-1][j]
// 			} else {
// 				dugeon[i][j] += min(dugeon[i-1][j], dugeon[i][j-1])
// 			}
// 		}
// 	}
// 	dp := make([][]int, len(dugeon))
// 	for i := 0; i < len(dugeon); i++ {
// 		dugeon[i] = make([]int, len(dugeon[i]))
// 	}
// 	for i := 0; i < len(dugeon); i++ {
// 		for j := 0; j < len(dugeon[i]); j++ {
// 			if i == 0 && j == 0 {
// 				if dugeon[0][0] < 0 {
// 					dp[0][0] = dugeon[0][0]
// 				}
// 				continue
//              // wtf what is the transfer equation ...
// 			}
// 		}
// 	}

// 	return dugeon[len(dugeon)-1][len(dugeon[0])-1]
// }
var resu, hp, current int

func calculateMinimumHP(dugeon [][]int) int {
	current = 0
	hp += dugeon[0][0]
	dfs(dugeon, 0, 0)
	return resu + 1
}

func dfs(dugeon [][]int, i, j int) {
	if hp < 0 && -hp > current {
		current = hp
	}
	if i == len(dugeon)-1 && j == len(dugeon[0])-1 {
		if resu == 0 || current < resu {
			resu = -current
		}
		return
	}
	if i < len(dugeon)-1 {
		hp += dugeon[i+1][j]
		dfs(dugeon, i+1, j)
		hp -= dugeon[i+1][j]
	}
	if j < len(dugeon[i])-1 {
		hp += dugeon[i][j+1]
		dfs(dugeon, i, j+1)
		hp -= dugeon[i][j+1]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
