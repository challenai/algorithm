package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var ptr *ListNode
	ptr = head
	list := []int{}
	for ptr != nil {
		list = append(list, ptr.Val)
		ptr = ptr.Next
	}

	mid := (len(list) - 1) / 2
	root := &TreeNode{
		Val: list[mid],
	}

	buildSubTree(list, root, 0, mid-1)
	buildSubTree(list, root, mid+1, len(list)-1)

	return root
}

func buildSubTree(list []int, root *TreeNode, left, right int) {
	if left > right {
		return
	}
	mid := (left + right) / 2
	if root.Val < list[mid] {
		root.Right = &TreeNode{
			Val: list[mid],
		}
		buildSubTree(list, root.Right, left, mid-1)
		buildSubTree(list, root.Right, mid+1, right)
	} else {
		root.Left = &TreeNode{
			Val: list[mid],
		}
		buildSubTree(list, root.Left, left, mid-1)
		buildSubTree(list, root.Left, mid+1, right)
	}
}
