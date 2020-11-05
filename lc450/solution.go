package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var ptr, par *TreeNode
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			root = root.Right
			return root
		}
		if root.Right == nil {
			root = root.Left
			return root
		}
		ptr = root.Right
		if ptr.Left == nil {
			root.Right = root.Right.Right
			return root
		}
		for ptr.Left != nil {
			ptr = ptr.Left
		}
		root.Val = ptr.Left.Val
		ptr.Left = nil
		return root
	}
	ptr = root
	for ptr != nil {
		if key == ptr.Val {
			break
		} else if key > ptr.Val {
			par = ptr
			ptr = ptr.Right
		} else {
			par = ptr
			ptr = ptr.Left
		}
	}
	if ptr == nil {
		return root
	}
	if ptr.Left == nil && ptr.Right == nil {
		if ptr.Val > par.Val {
			par.Right = nil
		} else {
			par.Left = nil
		}
		return root
	} else if ptr.Left == nil || ptr.Right == nil {
		if ptr.Left == nil {
			ptr.Val = ptr.Right.Val
			ptr.Left = ptr.Right.Left
			ptr.Right = ptr.Right.Right
		} else {
			ptr.Val = ptr.Left.Val
			ptr.Right = ptr.Left.Right
			ptr.Left = ptr.Left.Left
		}
	} else {
		if ptr.Right.Left == nil {
			ptr.Val = ptr.Right.Val
			ptr.Right = ptr.Right.Right
			return root
		}
		par = ptr
		ptr = ptr.Right
		for ptr.Left != nil {
			ptr = ptr.Left
		}
		par.Val = ptr.Left.Val
		ptr.Left = nil
	}
	return root
}
