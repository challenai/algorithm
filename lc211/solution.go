package solution

type node struct {
	children map[byte]*node
	isEnd    bool
}

type WordDictionary struct {
	*node
	searchSucc bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		node: &node{
			children: map[byte]*node{},
		},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	ptr := this.node
	for i := 0; i < len(word); i++ {
		if _, ok := ptr.children[word[i]]; !ok {
			ptr.children[word[i]] = &node{
				children: map[byte]*node{},
			}
		}
		ptr = ptr.children[word[i]]
	}
	ptr.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	this.searchSucc = false
	this.search(word, 0, this.node)
	println(this.searchSucc)
	return this.searchSucc
}

func (this *WordDictionary) search(word string, startIdx int, n *node) {
	if this.searchSucc {
		return
	}
	ptr := n
	for i := startIdx; i < len(word); i++ {
		if word[i] == '.' {
			if i == len(word)-1 {
				this.searchSucc = true
				return
			}
			for _, n := range ptr.children {
				this.search(word, i+1, n)
			}
		} else if _, ok := ptr.children[word[i]]; !ok {
			return
		} else {
			ptr = ptr.children[word[i]]
		}
	}
	if ptr.isEnd {
		this.searchSucc = true
	}
}
