package solution

import (
	"testing"
)

func TestMaxPrices(t *testing.T) {
	prices1 := []int{2, 4, 1}
	k1 := 2
	r1 := 2
	prices2 := []int{3, 2, 6, 5, 0, 3}
	k2 := 2
	r2 := 7

	resu1 := maxProfit(prices1, k1)
	resu2 := maxProfit(prices2, k2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
