package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resuList []*TreeNode

func flatten(root *TreeNode) {
	// I have no idea on how to realize an O(1) sc algorithm, I consider it exists a rotate node method, but rotate condition is a bit hard to figure out
	// and maybe I can use preorder to line up all the elements
	// :) have a try with preorder traversal
	if root == nil {
		return
	}
	resuList = []*TreeNode{}
	preorderTraversal(root)
	for i := 0; i < len(resuList)-1; i++ {
		resuList[i].Left = nil
		resuList[i].Right = resuList[i+1]
	}
	// var ptr *TreeNode
	// ptr = resuList[0]
	// for ptr != nil {
	// 	print(ptr.Val, " ")
	// 	ptr = ptr.Right
	// }
}

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	resuList = append(resuList, root)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}
