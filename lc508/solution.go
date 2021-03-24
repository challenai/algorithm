package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func FindFrequentTreeSum(root *TreeNode) []int {
	resu := []int{}
	hash := map[int]int{}
	loop(root, hash)
	for k, v := range hash {
		println(k, v)
		if len(resu) == 0 {
			resu = append(resu, k)
			continue
		}
		if v > hash[resu[0]] {
			resu[0] = k
		}
	}
	for k, v := range hash {
		if v == hash[resu[0]] && resu[0] != k {
			resu = append(resu, k)
		}
	}
	return resu
}

func loop(root *TreeNode, hash map[int]int) int {
	if root == nil {
		return 0
	}
	current := loop(root.Left, hash) + loop(root.Right, hash) + root.Val
	if _, ok := hash[current]; !ok {
		hash[current] = 1
	} else {
		hash[current]++
	}
	return current
}
