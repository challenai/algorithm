package solution

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	highestIdx := 0
	for i := 0; i < len(height); i++ {
		if height[i] > height[highestIdx] {
			highestIdx = i
		}
	}
	left := 0
	right := len(height) - 1
	sum := 0
	prevIdx := left
	for left < highestIdx {
		if height[left] < height[prevIdx] {
			sum += height[prevIdx] - height[left]
		} else if height[left] > height[prevIdx] {
			prevIdx = left
		}
		left++
	}
	prevIdx = right
	for highestIdx < right {
		if height[right] < height[prevIdx] {
			sum += height[left] - height[prevIdx]
		} else if height[right] > height[prevIdx] {
			prevIdx = right
		}
		right--
	}
	return sum
}
