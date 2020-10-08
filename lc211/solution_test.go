package solution

import (
	"testing"
)

func TestWordDict(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	wd.AddWord("madfd434t1r3f")
	if wd.Search("pad") {
		t.Fail()
	}
	if !wd.Search("bad") {
		t.Fail()
	}
	if !wd.Search(".ad") {
		t.Fail()
	}
	if !wd.Search("b..") {
		t.Fail()
	}
	if !wd.Search("b.d") {
		t.Fail()
	}
	if wd.Search(".b.") {
		t.Fail()
	}
	// how to handle this crazy one..
	if !wd.Search(".....43.t.r3.") {
		t.Fail()
	}
}
