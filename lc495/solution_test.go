package solution

import (
	"testing"
)

func TestFindPosisonedDuration(t *testing.T) {

	timeSeries1 := []int{1, 4}
	duration1 := 2
	r1 := 4
	timeSeries2 := []int{1, 2, 3, 4, 6, 9}
	duration2 := 2
	r2 := 9

	rsu1 := FindPoisonedDuration(timeSeries1, duration1)
	rsu2 := FindPoisonedDuration(timeSeries2, duration2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
