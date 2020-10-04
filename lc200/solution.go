package solution

import "strconv"

func numIslands(grid [][]byte) int {
	dsu := [][][2]int{}
	rsuHash := map[string]bool{}
	for i := 0; i < len(grid); i++ {
		dsu = append(dsu, make([][2]int, len(grid[i])))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			dsu[i][j] = [2]int{-1, -1}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				if i == 0 && j == 0 {
					dsu[i][j][0] = i
					dsu[i][j][1] = j
				} else if i == 0 {
					if grid[i][j-1] == '1' {
						dsu[i][j] = dsu[i][j-1]
					} else {
						dsu[i][j][0] = i
						dsu[i][j][1] = j
					}
				} else if j == 0 {
					if grid[i-1][j] == '1' {
						dsu[i][j] = dsu[i-1][j]
					} else {
						dsu[i][j][0] = i
						dsu[i][j][1] = j
					}
				} else {
					if grid[i-1][j] == '1' && grid[i][j-1] == '1' {
						// grid
						dsu[i][j] = dsu[i-1][j]
						temp := dsu[i][j-1]
						dsu[temp[0]][temp[1]] = dsu[i][j]
					} else if grid[i-1][j] == '1' {
						dsu[i][j] = dsu[i-1][j]
					} else if grid[i][j-1] == '1' {
						dsu[i][j] = dsu[i][j-1]
					} else {
						dsu[i][j][0] = i
						dsu[i][j][1] = j
					}
				}
			}
		}
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			findRoot(dsu, i, j)
			key := serialize(dsu[i][j][0], dsu[i][j][1])
			if _, ok := rsuHash[key]; !ok {
				rsuHash[key] = false
			}
		}
	}

	return len(rsuHash) - 1
}

func serialize(a, b int) string {
	return strconv.Itoa(a) + "_" + strconv.Itoa(b)
}

func findRoot(dsu [][][2]int, i, j int) {
	if dsu[i][j][0] == -1 || dsu[i][j][1] == -1 {
		return
	}
	var i_, j_ int
	i_ = i
	j_ = j
	for dsu[i_][j_][0] != i_ || dsu[i_][j_][1] != j_ {
		i_ = dsu[i_][j_][0]
		j_ = dsu[i_][j_][1]
	}
	dsu[i][j][0] = i_
	dsu[i][j][1] = j_
}
