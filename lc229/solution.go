package solution

func majorityElement(nums []int) []int {
	rsu := []int{}
	v1 := -1
	v2 := -1
	c1 := 0
	c2 := 0
	for i := 0; i < len(nums); i++ {
		if v1 != -1 && nums[i] == nums[v1] {
			c1++
			continue
		}
		if v2 != -1 && nums[i] == nums[v2] {
			c2++
			continue
		}
		if v1 == -1 {
			v1 = i
			c1 = 1
			continue
		}
		if v2 == -1 {
			v2 = i
			c2 = 1
			continue
		}
		c1--
		c2--
		if c1 == 0 {
			v1 = -1
		}
		if c2 == 0 {
			v2 = -1
		}
	}
	if v1 != -1 {
		rsu = append(rsu, nums[v1])
	}
	if v2 != -1 {
		rsu = append(rsu, nums[v2])
	}
	return rsu
}
