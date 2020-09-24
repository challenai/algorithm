package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var nums_ []int

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nums_ = nums
	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	installLeftTree(root, 0, mid-1)
	installRightTree(root, mid+1, len(nums)-1)
	return root
}

func installLeftTree(root *TreeNode, left, right int) {
	if left > right {
		return
	}
	mid := (left + right) / 2
	root.Left = &TreeNode{
		Val: nums_[mid],
	}
	installLeftTree(root.Left, left, mid-1)
	installRightTree(root.Left, mid+1, right)
}

func installRightTree(root *TreeNode, left, right int) {
	if left > right {
		return
	}
	mid := (left + right) / 2
	root.Right = &TreeNode{
		Val: nums_[mid],
	}
	installLeftTree(root.Right, left, mid-1)
	installRightTree(root.Right, mid+1, right)
}
