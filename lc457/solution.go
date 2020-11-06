package solution

func circularArrayLoop(nums []int) bool {
	// I catch up a O(n) tc algorithm, but need O(n) space ...
	// I found it wrong, when use a dsu to store the last el..
	// so just use brute force to handle this problem, and we got an O(n^2) tc algorithm, but with only O(1) space lol

	var ptr int
	var jump int
	for i := 0; i < len(nums); i++ {
		ptr = i
		jump = 0
		for nums[ptr]*nums[i] > 0 {
			jump++
			ptr += nums[ptr]
			if ptr >= len(nums) {
				ptr %= len(nums)
			} else if ptr < 0 {
				ptr += len(nums)
			}
			if ptr == i {
				if jump > 1 {
					return true
				}
				break
			}
		}
	}
	return false
}
