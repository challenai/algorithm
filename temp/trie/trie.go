package trie

type node struct {
	v        byte
	children map[byte]*node
	end      bool
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{
			children: map[byte]*node{},
		},
	}
}

func (t *Trie) Insert(s string) {
	// ptr := t.root
}

func (t *Trie) Search(s string) bool {
	return false
}

func (t *Trie) SearchPrefix(s string) []string {
	rsu := []int{}
	return rsu
}
