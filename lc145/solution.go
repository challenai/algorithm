package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resu []int

func postorderTraversal(root *TreeNode) []int {
	resu = []int{}
	dfs(root)
	return resu
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	resu = append(resu, root.Val)
}
