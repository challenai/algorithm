package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resu [][]int
var current []int

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return resu
	}
	if sum < 0 {
		return resu
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			temp := make([]int, len(current))
			copy(temp, current)
			resu = append(resu, temp)
		}
		return resu
	}
	current = append(current, root.Val)
	pathSum(root.Left, sum-root.Val)
	current = current[:len(current)-1]
	current = append(current, root.Val)
	pathSum(root.Right, sum-root.Val)
	current = current[:len(current)-1]

	return resu
}

func init() {
	resu = [][]int{}
	current = []int{}
}
