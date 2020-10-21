package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return max(rob(root.Left)+rob(root.Right), nonRob(root.Left)+nonRob(root.Right)+root.Val)
}

func nonRob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return rob(root.Left) + rob(root.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
