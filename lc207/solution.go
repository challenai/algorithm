package solution

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	deps := map[int][]int{}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		deps[i] = []int{}
	}
	for i := 0; i < len(prerequisites); i++ {
		if prerequisites[i][0] < numCourses && prerequisites[i][0] >= 0 {
			indegree[prerequisites[i][0]]++
			deps[prerequisites[i][1]] = append(deps[prerequisites[i][1]], prerequisites[i][0])
		}
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		temp, _ := deps[q[0]]
		delete(deps, q[0])
		q = q[1:]
		for i := 0; i < len(temp); i++ {
			indegree[temp[i]]--
			if indegree[temp[i]] == 0 {
				q = append(q, temp[i])
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] > 0 {
			return false
		}
	}
	return true
}
