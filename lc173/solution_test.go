package solution

import (
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 8,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	temp := 0
	// 1 2 5 6 7 8 9
	iterator := Constructor(root)
	iterator.Next()        // 1
	temp = iterator.Next() // 2
	if temp != 2 {
		t.Fail()
	}
	iterator.HasNext()     // return true
	iterator.Next()        // return 5
	iterator.HasNext()     // return true
	temp = iterator.Next() // return 6
	if temp != 6 {
		t.Fail()
	}
	iterator.HasNext() // return true
	iterator.Next()    // return 7
	iterator.Next()
	iterator.Next()
	if !iterator.HasNext() {
		t.Fail()
	}
}
