package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var rsu, maxCount, counter, currentNum int

func findMode(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxCount = 0
	counter = 0
	dfs(root)

	return rsu
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	if counter > 0 && root.Val == currentNum {
		counter++
		if counter > maxCount {
			rsu = root.Val
			maxCount = counter
		}
	} else {
		currentNum = root.Val
		counter = 1
	}
	dfs(root.Right)
}
