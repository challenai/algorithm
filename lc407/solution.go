package solution

import (
	"strconv"
)

var path map[string]bool
var pathReal [][]int
var heights_ [][]int
var currentHeightLevel int
var lowestBorder int
var rsu int
var flowSucc bool

func trapRainWater(heights [][]int) int {
	if len(heights) < 3 || len(heights[0]) < 3 {
		return 0
	}
	heights_ = heights
	rsu = 0

	for i := 1; i < len(heights_)-1; i++ {
		for j := 1; j < len(heights_[i])-1; j++ {
			if heights_[i-1][j] > heights_[i][j] && heights_[i][j-1] > heights_[i][j] {
				path = map[string]bool{}
				pathReal = [][]int{}
				flowSucc = true
				currentHeightLevel = heights_[i][j]
				lowestBorder = min(heights_[i-1][j], heights_[i][j-1])
				bfs(i, j)
				// println(flowSucc, lowestBorder, i, j, heights_[i][j])
				if flowSucc {
					for i := 0; i < len(pathReal); i++ {
						// print(pathReal[i][0], " ", pathReal[i][1], ";")
						// println(lowestBorder, "!")
						rsu += (lowestBorder - heights_[pathReal[i][0]][pathReal[i][1]])
						heights_[pathReal[i][0]][pathReal[i][1]] = lowestBorder
					}
				}
				flowSucc = true
			}
		}
	}
	return rsu
}

func bfs(i, j int) {
	if _, ok := path[serialize(i, j)]; ok {
		return
	}
	path[serialize(i, j)] = false
	if i == 0 || j == 0 || i == len(heights_)-1 || j == len(heights_[0])-1 {
		if heights_[i][j] <= currentHeightLevel {
			flowSucc = false
		} else {
			lowestBorder = min(lowestBorder, heights_[i][j])
		}
		return
	}
	pathReal = append(pathReal, []int{i, j})
	temp := [4][]int{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
	for k := 0; k < len(temp); k++ {
		if heights_[i][j] >= heights_[temp[k][0]][temp[k][1]] {
			bfs(temp[k][0], temp[k][1])
		} else if heights_[temp[k][0]][temp[k][1]] > currentHeightLevel {
			lowestBorder = min(lowestBorder, heights_[temp[k][0]][temp[k][1]])
		}
	}
}

func serialize(i, j int) string {
	return strconv.Itoa(i) + "_" + strconv.Itoa(j)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
