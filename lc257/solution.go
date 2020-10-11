package solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var current []string
var resu []string

func binaryTreePaths(root *TreeNode) []string {
	current = []string{}
	resu = []string{}
	dfs(root)
	return resu
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	current = append(current, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		resu = append(resu, strings.Join(current, "->"))
	}
	dfs(root.Left)
	dfs(root.Right)
	current = current[:len(current)-1]
}
