package solution

import (
	"strings"
)

type node struct {
	children map[byte]*node
	isEnd    bool
}

type trie struct {
	*node
}

func newTrie() *trie {
	return &trie{
		node: &node{
			children: map[byte]*node{},
		},
	}
}

func (t *trie) add(s string) {
	ptr := t.node
	for i := 0; i < len(s); i++ {
		if _, ok := ptr.children[s[i]]; !ok {
			ptr.children[s[i]] = &node{
				children: map[byte]*node{},
				isEnd:    false,
			}
		}
		ptr = ptr.children[s[i]]
	}
	ptr.isEnd = true
}

func (t *trie) search(s string) bool {
	ptr := t.node
	for i := 0; i < len(s); i++ {
		if _, ok := ptr.children[s[i]]; ok {
			ptr = ptr.children[s[i]]
		} else {
			return false
		}
	}
	return ptr.isEnd
}

func (t *trie) searchAllPotiential(s string) []int {
	resu := []int{}
	counter := 0
	ptr := t.node
	for i := 0; i < len(s); i++ {
		if _, ok := ptr.children[s[i]]; ok {
			ptr = ptr.children[s[i]]
			if ptr.isEnd {
				resu = append(resu, counter)
			}
			counter++
		} else {
			return resu
		}
	}
	return resu
}

var t *trie
var resu []string
var s_ string
var current []string

func wordBreak(s string, wordDict []string) []string {
	resu = []string{}
	s_ = s
	t = newTrie()
	current = []string{}
	for i := 0; i < len(wordDict); i++ {
		t.add(wordDict[i])
	}
	dfs(0)
	// for i := 0; i < len(resu); i++ {
	// 	println(resu[i])
	// }
	return resu
}

func dfs(startIdx int) {
	if startIdx == len(s_) {
		resu = append(resu, strings.Join(current, " "))
		return
	}
	potientialWords := t.searchAllPotiential(s_[startIdx:])
	for i := 0; i < len(potientialWords); i++ {
		current = append(current, s_[startIdx:startIdx+potientialWords[i]+1])
		startIdx += potientialWords[i] + 1
		dfs(startIdx)
		startIdx -= potientialWords[i] + 1
		current = current[:len(current)-1]
	}
}
