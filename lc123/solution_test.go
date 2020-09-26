package solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	r1 := 6
	prices2 := []int{1, 2, 3, 4, 5}
	r2 := 4
	prices3 := []int{7, 6, 4, 3, 1}
	r3 := 0
	prices4 := []int{1}
	r4 := 0

	resu1 := maxProfit(prices1)
	resu2 := maxProfit(prices2)
	resu3 := maxProfit(prices3)
	resu4 := maxProfit(prices4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
