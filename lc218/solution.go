package solution

func getSkyline(buildings [][]int) [][]int {
	// 每来一个新的建筑
	// 比较左端点和 已有建筑群的最小右端点，如果比最小的小，则连起来，否则最新小的被踢出去
	// 如果左端点小有交集了，则比高度，如果高，则有新天际线生成。。。

	// 我们把交集按右端点大小升序，从而获得O(1)的 最小右端点搜索速度。
	// 问题在于插入变成了O(n), 使用二分搜索可以降低到logn
	// 这让我想到了二叉树和单调栈，小二叉树显然可以用, 但是实现成本略高
	// 那么单调栈呢？单调栈可以提供O(1)查找和插入。
	// 似乎是不行的，因为我们还要考虑左端点

	// shit. 我发现右端点和高度都是要获取最小的
	// 崩了。。。
	intersect := [][]int{}
	resu := [][]int{}
	for i := 0; i < len(buildings); i++ {
		if len(intersect) == 0 {
			resu = append(resu, []int{buildings[i][0], buildings[i][2]})
			intersect = append(intersect, buildings[i])
			continue
		}
		if buildings[i][0] <= intersect[0][1] {
			if buildings[i][2] != intersect[i][2] {
				resu = append(resu, []int{buildings[i][1], buildings[i][2]})
			}
			intersect = append(intersect, buildings[i])
		} else {

		}
	}
}
