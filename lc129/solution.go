package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var current int
var resu int

func sumNumbers(root *TreeNode) int {
	current = 0
	resu = 0
	loop(root)
	return resu
}

func loop(root *TreeNode) {
	if root == nil {
		return
	}
	current *= 10
	current += root.Val
	if root.Left == nil && root.Right == nil {
		resu += current
	}
	loop(root.Left)
	loop(root.Right)
	current /= 10
}
