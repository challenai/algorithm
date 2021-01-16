package main

type node struct {
	v, cnt, dulp int
	left, right  *node
}

type Bst struct {
	root *node
}

func NewTree() *Bst {
	return &Bst{
		root: nil,
	}
}

func (tr *Bst) Insert(v int) {
	newNode := &node{
		v:    v,
		dulp: 1,
		cnt:  1,
	}
	if tr.root == nil {
		tr.root = newNode
		return
	}
	ptr := tr.root
	var par *node
	for ptr != nil {
		if v == ptr.v {
			ptr.dulp++
			ptr.cnt++
			return
		}
		ptr.cnt++
		par = ptr
		if v > ptr.v {
			ptr = ptr.right
		} else {
			ptr = ptr.left
		}
	}
	if v > par.v {
		par.right = newNode
	} else {
		par.left = newNode
	}
}

func (tr *Bst) Remove(v int) {
	if tr.root == nil {
		return
	}
	if tr.root.v == v {
		// if tr.root.cnt == 1
	}
	// ptr := tr.root
	// var par *node
	// for ptr != nil {
	// 	if ptr.
	// }
}

func (tr *Bst) FindKth(k int) int {
	if tr.root == nil {
		return -1
	}
	return -1
}
