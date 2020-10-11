package solution

func hIndex2(citations []int) int {
	resu := 0
	for resu < len(citations) {
		if citations[len(citations)-1-resu] < resu+1 {
			break
		}
		resu++
	}
	return resu
}
