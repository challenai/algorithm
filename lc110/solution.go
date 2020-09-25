package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(countDeepth(root.Left)-countDeepth(root.Right)) <= 1
}

func countDeepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(countDeepth(root.Left), countDeepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
