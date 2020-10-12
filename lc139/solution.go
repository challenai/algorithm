package solution

var t *Trie
var exist bool
var s_ string

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	s_ = s
	t = NewTrie()
	exist = false
	for i := 0; i < len(wordDict); i++ {
		t.insert(wordDict[i])
	}

	dfs(0)
	return exist
}

func dfs(startIdx int) {
	if exist {
		return
	}
	if startIdx >= len(s_) {
		exist = true
		return
	}
	currentPontiential := t.SearchPrefix(s_[startIdx:])
	for i := 0; i < len(currentPontiential); i++ {
		dfs(startIdx + currentPontiential[i] + 1)
	}
}

type Trie struct {
	*node
}

type node struct {
	children map[byte]*node
	isEnd    bool
}

func NewTrie() *Trie {
	return &Trie{
		node: &node{
			children: map[byte]*node{},
		},
	}
}

func (t *Trie) SearchPrefix(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	resu := []int{}
	ptr := t.node
	for i := 0; i < len(s); i++ {
		if _, ok := ptr.children[s[i]]; ok {
			ptr = ptr.children[s[i]]
			if ptr.isEnd {
				resu = append(resu, i)
			}
		} else {
			break
		}
	}
	return resu
}

func (t *Trie) insert(s string) {
	if len(s) == 0 {
		return
	}
	ptr := t.node
	for i := 0; i < len(s); i++ {
		if _, ok := ptr.children[s[i]]; !ok {
			ptr.children[s[i]] = &node{
				children: map[byte]*node{},
			}
		}
		ptr = ptr.children[s[i]]
	}
	ptr.isEnd = true
}
