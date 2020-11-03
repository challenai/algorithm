package solution

import (
	"testing"
)

func TestAllOne(t *testing.T) {
	allOne := Constructor()

	allOne.Inc("apple")
	allOne.Inc("apple")
	allOne.Inc("apple")

	allOne.Inc("banana")
	allOne.Inc("banana")
	// allOne.Dec("banana")

	allOne.Inc("trump")

	allOne.Inc("lemon")
	allOne.Inc("lemon")
	allOne.Inc("lemon")
	allOne.Inc("lemon")

	if allOne.GetMaxKey() != "apple" {
		t.Fail()
	}
	println(allOne.GetMaxKey())
	println(allOne.GetMinKey())
	// if allOne.GetMinKey() != "banana" {
	// 	t.Fail()
	// }
	t.Fail()
}
