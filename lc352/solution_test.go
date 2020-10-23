package solution

import (
	"testing"
)

func TestDisjointStream(t *testing.T) {
	stream := []int{1, 3, 7, 2, 6, 5, 12, 4, 9, 11}
	r := [][]int{{1, 7}, {9, 9}, {11, 12}}
	st := Constructor()
	for i := 0; i < len(stream); i++ {
		st.AddNum(stream[i])
	}
	rsu := st.GetIntervals()
	if len(r) != len(rsu) {
		t.Fail()
	}
	for i := 0; i < len(r); i++ {
		if r[i][0] != rsu[i][0] || r[i][1] != rsu[i][1] {
			t.Fail()
		}
	}
}
