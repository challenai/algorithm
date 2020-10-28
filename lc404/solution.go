package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root.Left, true) + dfs(root.Right, false)
}

func dfs(root *TreeNode, isLeftChild bool) int {
	if root == nil {
		return 0
	}
	if isLeftChild && root.Left == nil && root.Right == nil {
		return root.Val
	}
	return dfs(root.Left, true) + dfs(root.Right, false)
}
