package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var path [2][]*TreeNode
var current []*TreeNode
var resu *TreeNode
var currentSeq int
var isEnd bool

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == q {
		return p
	}
	path = [2][]*TreeNode{}
	current = []*TreeNode{}
	resu = root
	isEnd = false
	currentSeq = 0
	dfs(root, p)
	isEnd = false
	currentSeq = 1
	dfs(root, q)
	for i := 0; i < min(len(path[0]), len(path[1])); i++ {
		if path[0][i] == path[1][i] {
			resu = path[0][i]
		}
	}
	return resu
}

func dfs(root, n *TreeNode) {
	if root == nil || isEnd {
		return
	}
	if root == n {
		temp := make([]*TreeNode, len(current))
		copy(temp, current)
		path[currentSeq] = temp
		isEnd = true
	}

	current = append(current, root)
	dfs(root.Left, n)
	dfs(root.Right, n)
	current = current[:len(current)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
