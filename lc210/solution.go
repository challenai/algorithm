package solution

func findOrder(nums int, pre [][]int) []int {
	resu := []int{}
	indegree := make([]int, nums)
	q := []int{}
	for i := 0; i < len(pre); i++ {
		if pre[i][1] >= 0 && pre[i][1] < nums {
			indegree[pre[i][1]]++
		}
	}
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			q = append(q, indegree[i])
		}
	}
	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		for i := 0; i < len(pre); i++ {
			if pre[i][1] == temp {

			}
		}
	}
}
