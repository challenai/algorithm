package solution

var current, longestPath []int
var edgesHash map[int][]int
var existNodes map[int]bool

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	rsu := []int{}
	// dfs, path record, find mid in list
	existNodes = map[int]bool{}
	edgesHash = map[int][]int{}
	longestPath = []int{}
	for i := 0; i < n; i++ {
		edgesHash[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		edgesHash[edges[i][0]] = append(edgesHash[edges[i][0]], edges[i][1])
		edgesHash[edges[i][1]] = append(edgesHash[edges[i][1]], edges[i][0])
	}
	for k, v := range edgesHash {
		if len(v) == 1 {
			// existNodes[k] = false
			// current = append(current, k)
			dfs(k)
			break
		}
	}

	newStartPoint := longestPath[len(longestPath)-1]
	dfs(newStartPoint)
	rsu = append(rsu, longestPath[len(longestPath)>>1])
	// for i := 0; i < len(longestPath); i++ {
	// 	print(longestPath[i], " ")
	// }
	// println()
	if len(longestPath)%2 == 0 {
		rsu = append(rsu, longestPath[((len(longestPath)-1)>>1)+1])
	}

	return rsu
}

func dfs(startNode int) {
	existNodes[startNode] = false
	current = append(current, startNode)
	if temp, ok := edgesHash[startNode]; ok {
		if len(temp) == 1 && len(current) > len(longestPath) {
			longestPath = make([]int, len(current))
			copy(longestPath, current)
		}
	}
	potientialNodes, _ := edgesHash[startNode]
	for i := 0; i < len(potientialNodes); i++ {
		if _, ok := existNodes[potientialNodes[i]]; !ok {
			dfs(potientialNodes[i])
		}
	}
	current = current[:len(current)-1]
	delete(existNodes, startNode)
}
