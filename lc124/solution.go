package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var currentSum, maximum int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	currentSum = 0
	maximum = root.Val
	loop(root)
	return maximum
}

// func loop(root *TreeNode) int {
// 	if root == nil {
// return 0
// }
// leftPathSum := loop(root.Left)
// rightPathSum := loop(root.Right)
// currentSum = max(leftPathSum, rightPathSum)
// maximum = maxList([]int{maximum, leftPathSum, rightPathSum, leftPathSum + rightPathSum + root.Val})

// return max(leftPathSum, rightPathSum)
// }

func loop(root *TreeNode) {
	if root == nil {
		return
	}
	currentSum = 0
	loop(root.Left)
	currentSum = max(currentSum+root.Val, 0)
	maximum = max(maximum, currentSum)
	temp := currentSum
	currentSum = 0
	loop(root.Right)
	maximum = maxList([]int{maximum, currentSum + temp, currentSum + root.Val})
	currentSum = max(currentSum+root.Val, temp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxList(list []int) int {
	if len(list) == 0 {
		return 0
	}
	for i := 1; i < len(list); i++ {
		if list[i] > list[0] {
			list[0] = list[i]
		}
	}
	return list[0]
}
