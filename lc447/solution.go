package solution

func numberOfBoomerangs(points [][]int) int {
	hash := map[int][][]int{}
	var distanceSqr, rsu int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distanceSqr = sqr(points[i][0]-points[j][0]) + sqr(points[i][1]-points[j][1])
			if v, ok := hash[distanceSqr]; ok {
				for k := 0; k < len(v); k++ {
					if v[k][0] == i || v[k][1] == i || v[k][0] == j || v[k][1] == j {
						rsu += 2
					}
				}
				hash[distanceSqr] = append(hash[distanceSqr], []int{i, j})
			} else {
				hash[distanceSqr] = [][]int{[]int{i, j}}
			}
		}
	}
	return rsu
}

func sqr(a int) int {
	return a * a
}
