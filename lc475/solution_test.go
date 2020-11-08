package solution

import (
	"testing"
)

func TestFindRadius(t *testing.T) {
	house1 := []int{1, 2, 3}
	heaters1 := []int{2}
	r1 := 1
	house2 := []int{1, 2, 3, 4, 5, 6}
	heaters2 := []int{1, 4}
	r2 := 2
	house3 := []int{1, 5}
	heaters3 := []int{2}
	r3 := 3

	rsu1 := findRadius(house1, heaters1)
	rsu2 := findRadius(house2, heaters2)
	rsu3 := findRadius(house3, heaters3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
