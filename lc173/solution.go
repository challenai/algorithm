package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	st  []*TreeNode
	ptr *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	i := BSTIterator{
		st:  []*TreeNode{},
		ptr: root,
	}
	return i
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	resu := 0
	for this.ptr != nil {
		this.st = append(this.st, this.ptr)
		this.ptr = this.ptr.Left
	}
	if len(this.st) > 0 {
		this.ptr = this.st[len(this.st)-1]
		this.st = this.st[:len(this.st)-1]
		resu = this.ptr.Val
		this.ptr = this.ptr.Right
	}

	return resu
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.ptr == nil && len(this.st) == 0
}
