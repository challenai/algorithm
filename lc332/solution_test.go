package solution

import (
	"testing"
)

func TestX(t *testing.T) {
	ticket1 := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	r1 := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	ticket2 := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	r2 := []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	resu1 := findItinerary(ticket1)
	resu2 := findItinerary(ticket2)
	if len(r1) != len(resu1) {
		t.Fail()
	}
	if len(r2) != len(resu2) {
		t.Fail()
	}
	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1[i] {
			t.Fail()
		}
	}
	for i := 0; i < len(r2); i++ {
		if r2[i] != resu2[i] {
			t.Fail()
		}
	}

	if false {
		t.Fail()
	}
}
