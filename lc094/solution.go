package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resu []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return resu
	}
	inorderTraversal(root.Left)
	resu = append(resu, root.Val)
	inorderTraversal(root.Right)
	return resu
}

func init() {
	resu = []int{}
}

func inorderTraversal2(root *TreeNode) []int {
	st := []*TreeNode{}
	var ptr *TreeNode
	ptr = root
	for ptr != nil || len(st) > 0 {
		if ptr != nil {
			st = append(st, ptr)
			ptr = ptr.Left
		} else {
			ptr = st[len(st)-1]
			st = st[:len(st)-1]
			resu = append(resu, ptr.Val)
			ptr = ptr.Right
		}
	}
	return resu
}
