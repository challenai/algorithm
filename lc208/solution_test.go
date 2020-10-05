package solution

import (
	"testing"
)

func TestTrie(t *testing.T) {
	word := "apple"
	word2 := "banana"
	word3 := "append"
	word4 := "apply"

	trie := Constructor()
	trie.Insert(word)
	trie.Insert(word2)
	trie.Insert(word3)
	trie.Insert(word4)

	existApple := trie.Search(word)
	existApp := trie.Search("app")
	prefixApp := trie.StartsWith("app")
	existBanana := trie.Search(word2)
	existApply := trie.Search(word4)

	if !existApple || existApp || !prefixApp || !existBanana || !existApply {
		t.Fail()
	}
}
