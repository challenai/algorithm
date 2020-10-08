package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func revertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	revertTree(root.Left)
	revertTree(root.Right)
	return root
}
