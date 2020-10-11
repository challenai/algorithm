package solution

import (
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
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
		Val: 3,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}

	r := []string{"521", "5837", "589"}
	resu := binaryTreePaths(root)
	// for i := 0; i < len(resu); i++ {
	// 	print(resu[i], " ")
	// }

	if len(resu) != len(r) {
		t.Fail()
	}
}
