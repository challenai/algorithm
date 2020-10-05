package solution

type node struct {
	children map[byte]*node
	isEnd    bool
}

type Trie struct {
	*node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node: &node{
			children: map[byte]*node{},
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return true
	}
	ptr := this.node
	for i := 0; i < len(word); i++ {
		if _, ok := ptr.children[word[i]]; ok {
			ptr = ptr.children[word[i]]
		} else {
			return false
		}
	}
	return ptr.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	ptr := this.node
	for i := 0; i < len(prefix); i++ {
		if _, ok := ptr.children[prefix[i]]; ok {
			ptr = ptr.children[prefix[i]]
		} else {
			return false
		}
	}
	return true
}
