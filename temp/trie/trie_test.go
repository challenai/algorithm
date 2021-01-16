package trie

import "testing"

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	tr.Insert("append")
	tr.Insert("application")
	tr.Insert("banana")
	tr.Insert("amplification")
	if tr.Search("apps") {
		t.Fail()
	}
	if !tr.Search("append") {
		t.Fail()
	}
	if !tr.Search("banana") {
		t.Fail()
	}
}
