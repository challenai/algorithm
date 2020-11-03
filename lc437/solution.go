package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var rsu, target int

func pathSum(root *TreeNode, sum int) int {
	// idea: get all the potiential sum to leaf or root
	// enumerate all the node combination (par and child), and use O(1) minus
	// tc is O(n^2)
	// we replace val to sum so that we can get an inplace algorithm to save memory,
	// if we dont need to consider the traversal comsumption,
	// we can get an O(1) space algorithm.
	// now have a try:
	if root == nil {
		return 0
	}

	target = sum

	sumPath(root.Left, root.Val)
	sumPath(root.Right, root.Val)

	comparePath(root)
	traversalChildrenNodes(root, 0)

	return rsu
}

func sumPath(root *TreeNode, pathSum int) {
	if root == nil {
		return
	}
	root.Val += pathSum
	sumPath(root.Left, root.Val)
	sumPath(root.Right, root.Val)
}

func comparePath(root *TreeNode) {
	if root == nil {
		return
	}
	traversalChildrenNodes(root.Left, root.Val)
	traversalChildrenNodes(root.Right, root.Val)
	comparePath(root.Left)
	comparePath(root.Right)
}

func traversalChildrenNodes(root *TreeNode, pathSum int) {
	if root == nil {
		return
	}
	if root.Val-pathSum == target {
		rsu++
	}
	traversalChildrenNodes(root.Left, pathSum)
	traversalChildrenNodes(root.Right, pathSum)
}
