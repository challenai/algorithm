package solution

// type node struct {
// 	val         int
// 	left, right *node
// }

// type bst struct {
// 	head *node
// 	cnt  int
// }

// func newTree() *bst {
// 	return &bst{
// 		head: nil,
// 		cnt:  0,
// 	}
// }

// // insert element, if already exist the same return false
// func (t *bst) insert(v int) bool {
// 	if t.head == nil {
// 		t.head = &node{
// 			val: v,
// 		}
// 		t.cnt = 1
// 		return true
// 	}

// 	ptr := t.head
// 	var par *node
// 	for ptr != nil {
// 		par = ptr
// 		if v == ptr.val {
// 			return false
// 		} else if v > ptr.val {
// 			ptr = ptr.right
// 		} else {
// 			ptr = ptr.left
// 		}
// 	}
// 	if v > par.val {
// 		par.right = &node{
// 			val: v,
// 		}
// 	} else {
// 		par.left = &node{
// 			val: v,
// 		}
// 	}
// 	t.cnt++
// 	return true
// }

// func (t *bst) delete(v int) {
// 	if t.head == nil {
// 		return
// 	}
// 	if t.head.val == v {
// 		// todo
// 		return
// 	}
// 	ptr := t.head
// 	var par *node
// 	for ptr != nil {
// 		if v == ptr.val {
// 			// delete
// 			if ptr.left == nil && ptr.right == nil {
// 				// todo
// 				par.xx = nil
// 			} else if ptr.left == nil {
// 				par.xx = ptr.right
// 			} else if ptr.right == nil {
// 				par.xx = ptr.left
// 			} else {
// 				// find prev , and substitude the find one
// 				// make me realize that, build a bst is not realistic in interview except remeber all the code...
// 				// if you think about the process, you need some time, and the problem is building a bst will
// 				// stop your communicaton with the others, for example. the interviewer
// 				// even worse, I still need realize a treap or splay to make the tree auto rebalance...
// 				// so the conclusion is dont try to build a bst to help you solve problem,
// 				// but you can still talk about the prossibility of adopting bst
// 				// by the way, most of the time in my work, I will use Red-black-tree, but I dont want to repeat myself to build another one
// 				temp := ptr.left
// 				for ptr.right != nil {
// 					ptr = ptr.right
// 				}
// 			}
// 		} else if v > ptr.val {
// 			par = ptr
// 			ptr = ptr.right
// 		} else {
// 			par = ptr
// 			ptr = ptr.left
// 		}
// 	}
// }

// func (t *bst) getMax() int {
// 	if t.head == nil {
// 		return 0
// 	}
// 	ptr := t.head
// 	for ptr.right != nil {
// 		ptr = ptr.right
// 	}
// 	return ptr.val
// }

// func (t *bst) getMin() int {
// 	if t.head == nil {
// 		return 0
// 	}
// 	ptr := t.head
// 	for ptr.left != nil {
// 		ptr = ptr.left
// 	}
// 	return ptr.val
// }

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	// idea: we use a sliding window method
	// data in window stored in BST
	// go has no native bst, but I cant catch up any other nlogn algorithm...

	// I finally change the desicion , use a sorted list replace bst,
	// so the new complexity is nk, because I need to move all the items when insert
	// since we can only get a O(nk) complexity, we can just use a navie search in the window
	for i := 0; i < len(nums)-k; i++ {
		for j := i + 1; j <= k; j++ {
			if (nums[i] - nums[j]) <= t {
				return true
			}
		}
	}
	return false
	// I just use 2 min to build this new algorithm...
	// the key bst solution in my opinion is bst can provide a O(logn) speed to delete and insert, and getMin, getMax
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
