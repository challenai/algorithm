package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resu []int

func preorderTraversal(root *TreeNode) []int {
	resu = []int{}
	dfs(root)
	return resu
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	resu = append(resu, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

// func preorderTraversal(root *TreeNode) []int {
// 	st := []*TreeNode{}
// 	rsu := []int{}
// 	ptr := root
// 	for ptr != nil || len(st) > 0 {
// 		if ptr != nil {
// 			rsu = append(rsu, ptr.Val)
// 			if ptr.Right != nil {
// 				st = append(st, ptr.Right)
// 			}
// 			if ptr.Left != nil {
// 				st = append(st, ptr.Left)
// 			}
// 		} else {
// 			ptr = st[len(st)-1]
// 			st := st[:len(st)-1]
// 			rsu = append(rsu, ptr.Val)
// 		}
// 	}
// }
