package solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	r := 5
	resu := maxProfit(prices)

	if resu != r {
		t.Fail()
	}
}
