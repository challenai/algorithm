package solution

func candy(nums []int) int {
	resu := 0
	candies := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			candies[i] = 1
			continue
		}
		if nums[i] > nums[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			resu += candies[i]
			continue
		}
		if nums[i] > nums[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
		resu += candies[i]
	}
	return resu
}
