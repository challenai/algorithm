package solution

func computeArea(a, b, c, d, e, f, g, h int) int {
	// idea: caculate the intersection in each matrix,
	// so that we can get the intersection in all diamension

	// why I can handle this problem so easily?
	// maybe the inspiration of diamension
	rec1 := [][]int{{a, c}, {b, d}}
	rec2 := [][]int{{e, g}, {f, h}}
	intersection0 := caculateIntersection(rec1[0], rec2[0])
	intersection1 := caculateIntersection(rec1[1], rec2[1])

	return (c-a)*(d-b) + (g-e)*(h-f) - intersection0*intersection1
}

func caculateIntersection(range0, range1 []int) int {
	if min(range0) >= max(range1) || min(range1) >= max(range0) {
		return 0
	}

	return min([]int{range0[1], range1[1]}) - max([]int{range0[0], range1[0]})
}

func max(range0 []int) int {
	if range0[0] > range0[1] {
		return range0[0]
	}
	return range0[1]
}

func min(range0 []int) int {
	if range0[0] < range0[1] {
		return range0[0]
	}
	return range0[1]
}
