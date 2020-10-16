package solution

var edgesHash map[int][]int
var current, longestPath []int

func findMinHeightTrees2(n int, edges [][]int) []int {
	for i := 0; i < n; i++ {
		edgesHash[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		edgesHash[edges[i][0]] = append(edgesHash[edges[i][0]], edges[i][1])
	}

}
