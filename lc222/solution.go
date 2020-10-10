package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}
	// according to the properties of BST
	// use binary search for last element can get a n! algorithm~
	resu := 0
	ptr := root
	height := 0
	tempHeight := 0
	for ptr != nil {
		height++
		ptr = ptr.Left
	}
	ptr = root
	for ptr != nil {
		tempHeight++
		ptr = ptr.Right
	}
	if tempHeight == height {
		return getCountsFromHeight(height)
	}
	ptr = root
	var explorePtr *TreeNode
	var exploreHeight int
	baseCnt := getCountsFromHeight(tempHeight)
	resu += baseCnt + 1
	for ptr != nil {
		explorePtr = ptr.Left
		exploreHeight = 1
		for explorePtr != nil {
			explorePtr = explorePtr.Right
			exploreHeight++
		}
		if exploreHeight == tempHeight {
			ptr = ptr.Right
			resu += baseCnt
		} else {
			ptr = ptr.Left
		}
		tempHeight--
	}
	println(resu)
	return resu
}

func getCountsFromHeight(n int) int {
	resu := 0
	current := 1
	for i := 0; i < n; i++ {
		resu += current
		current <<= 1
	}
	return resu
}
