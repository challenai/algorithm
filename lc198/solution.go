package solution

func rob(nums []int) int {
	var front, current, temp int
	current = nums[0]
	for i := 0; i < len(nums); i++ {
		temp = current
		if front+nums[i] > current {
			current = front + nums[i]
		}
		front = temp
	}
	return current
}
