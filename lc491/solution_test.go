package solution

import (
	"testing"
)

func TestFindSubSequences(t *testing.T) {

	nums := []int{4, 6, 7, 7}
	r := [][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}}
	rsu := FindSubSequences(nums)

	// println(len(rsu))
	// for i := 0; i < len(rsu); i++ {
	// 	for j := 0; j < len(rsu[i]); j++ {
	// 		print(rsu[i][j], " ")
	// 	}
	// 	println()
	// }
	if len(rsu) != len(r) {
		t.Fail()
	}
}
