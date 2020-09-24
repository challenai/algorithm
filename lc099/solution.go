package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var wrongNode []*TreeNode

func recoverTree(root *TreeNode) *TreeNode {
	// var last, par, ptr *TreeNode
	// st := []*TreeNode{}
	// ptr = root
	// for ptr != nil || len(st) > 0 {
	// 	if last == nil {
	// 		last = ptr
	// 	}
	// 	par = ptr
	// 	if ptr != nil {
	// 		st = append(st, ptr)
	// 		ptr = ptr.Left
	// 	} else {
	// 		ptr = st[len(st)-1]
	// 		st = st[:len(st)-1]
	// if par.Val > ptr.Val {

	// }
	// ptr = ptr.Right
	// }
	// }

	findWrongNodes(root.Left)
	findWrongNodes(root.Right)
	println(len(wrongNode))
	if len(wrongNode) > 3 {
		wrongNode[1].Val, wrongNode[3].Val = wrongNode[3].Val, wrongNode[1].Val
	} else if len(wrongNode) == 2 {
		wrongNode[0].Val, wrongNode[1].Val = wrongNode[1].Val, wrongNode[0].Val
	}

	return root
}

func findWrongNodes(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	if root.Left != nil && root.Left.Val > root.Val {
		wrongNode = append(wrongNode, root)
		wrongNode = append(wrongNode, root.Left)
	}
	if root.Right != nil && root.Right.Val < root.Val {
		wrongNode = append(wrongNode, root)
		wrongNode = append(wrongNode, root.Right)
	}
	findWrongNodes(root.Left)
	findWrongNodes(root.Right)
}
