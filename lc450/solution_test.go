package solution

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
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

	root = deleteNode(root, 3)
	if root.Left.Right != nil {
		t.Fail()
	}
	root = deleteNode(root, 2)
	if root.Left.Val != 1 || root.Left.Left != nil {
		t.Fail()
	}
	root = deleteNode(root, 8)
	if root.Right.Val == 8 {
		t.Fail()
	}
	// dfs(root)
}

// func dfs(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	println(root.Val)
// 	dfs(root.Left)
// 	dfs(root.Right)
// }
