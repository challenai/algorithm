package solution

import (
	"testing"
)

var balance bool

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	resu := sortedArrayToBST(nums)
	balance = true
	checkBalance(resu)

	if !balance {
		t.Fail()
	}
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if !balance {
		return 0
	}
	leftHeight := checkBalance(root.Left)
	rightHeight := checkBalance(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		balance = false
		return 0
	}
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
