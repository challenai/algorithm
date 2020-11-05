package solution

import (
	"testing"
)

func TestBinaryTreeEncoder(t *testing.T) {
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
		Val: 81,
	}
	root.Right.Left = &TreeNode{
		Val: 6,
	}
	root.Right.Left.Right = &TreeNode{
		Val: 723,
	}
	root.Right.Right = &TreeNode{
		Val: 9,
	}
	encoder := Constructor()
	decoder := Constructor()
	str := encoder.Serialize(root)
	tr := decoder.Deserialize(str)
	// dfs(tr)

	if tr.Val != 5 || tr.Right.Val != 81 {
		t.Fail()
	}
}

// func dfs(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	println(root.Val)
// 	dfs(root.Left)
// 	dfs(root.Right)
// }
