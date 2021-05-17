package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var resu []*TreeNode
var n_, counter int
var used []bool
var current *TreeNode

// func generateTrees(n int) []*TreeNode {
// 	n_ = n
// 	resu = []*TreeNode{}
// 	used = make([]bool, n)
// 	counter = 0
// 	for i := 1; i <= n; i++ {
// 		used[i-1] = true
// 		counter++
// 		current = &TreeNode{
// 			Val: i + 1,
// 		}
// 		loop()
// 		counter--
// 		used[i-1] = false
// 	}
// 	println(len(resu))
// 	return resu
// }

// func loop() {
// 	if counter == len(used) {
// 		resu = append(resu, current)
// 		return
// 	}
// 	for i := 0; i < len(used); i++ {
// 		if !used[i] && (i > 0 && (!used[i-1]) || (i < len(used)-1 && !used[len(used)-1])) {
// 			used[i] = true
// 			counter++
// 			insert(current, i+1)
// 			loop()
// 			remove(current, i+1)
// 			counter--
// 			used[i] = false
// 		} else {
// 			used[i] = true
// 			counter++
// 			insert(current, i+1)
// 			loop()
// 			remove(current, i+1)
// 			counter--
// 			used[i] = false
// 		}
// 	}
// }

// func insert(root *TreeNode, n int) {
// 	ptr := root
// 	var par *TreeNode
// 	for ptr != nil {
// 		par = ptr
// 		if n > ptr.Val {
// 			ptr = ptr.Right
// 		} else {
// 			ptr = ptr.Left
// 		}
// 	}
// 	if n > par.Val {
// 		par.Right = &TreeNode{
// 			Val: n,
// 		}
// 	} else {
// 		par.Left = &TreeNode{
// 			Val: n,
// 		}
// 	}
// }

// func remove(root *TreeNode, n int) {
// 	ptr := root
// 	var par *TreeNode
// 	for ptr != nil {
// 		par = ptr
// 		if n == ptr.Val {
// 			break
// 		} else if n > ptr.Val {
// 			ptr = ptr.Right
// 		} else {
// 			ptr = ptr.Left
// 		}
// 	}
// 	if par.Left != nil && par.Left.Val == n {
// 		par.Left = nil
// 	}
// 	if par.Right != nil && par.Right.Val == n {
// 		par.Right = nil
// 	}
// }

func generateTrees(n int) []*TreeNode {
	for i := 1; i < n; i++ {

	}
}
