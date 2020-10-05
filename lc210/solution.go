package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	rsu := []int{}
	indegree := make([]int, numCourses)
	nodes := map[int][]int{}
	for i := 0; i < numCourses; i++ {
		nodes[i] = []int{}
	}
	for i := 0; i < len(prerequisites); i++ {
		indegree[prerequisites[i][0]]++
		nodes[prerequisites[i][1]] = append(nodes[prerequisites[i][1]], prerequisites[i][0])
	}
	q := []int{}
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		temp, _ := nodes[q[0]]
		rsu = append(rsu, q[0])
		delete(nodes, q[0])
		q = q[1:]
		for i := 0; i < len(temp); i++ {
			indegree[temp[i]]--
			if indegree[temp[i]] == 0 {
				q = append(q, temp[i])
			}
		}
	}
	for i := 0; i < len(indegree); i++ {
		if indegree[i] > 0 {
			return []int{}
		}
	}

	return rsu
}
