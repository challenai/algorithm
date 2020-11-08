package solution

import (
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	r := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}

	rsu := findAllConcatenatedWordsInADict(words)

	if len(rsu) != len(r) {
		t.Fail()
	}
}
