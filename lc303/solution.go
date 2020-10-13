package solution

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	temp := NumArray{
		sums: make([]int, len(nums)),
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			temp.sums[i] = nums[i]
			continue
		}
		temp.sums[i] = temp.sums[i-1] + nums[i]
	}
	return temp
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}
