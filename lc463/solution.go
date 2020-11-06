package solution

func islandPerimeter(grid [][]int) int {
	rsu := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				rsu++
			}
			if i == len(grid)-1 || grid[i+1][j] == 0 {
				rsu++
			}
			if j == 0 || grid[i][j-1] == 0 {
				rsu++
			}
			if j == len(grid[i])-1 || grid[i][j+1] == 0 {
				rsu++
			}
		}
	}
	return rsu
}
