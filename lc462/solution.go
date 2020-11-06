package solution

import (
	"math/rand"
	// "sort"
)

func minMoves2(nums []int) int {
	rsu := 0
	// 寻找第k大的数
	num := quickSelectMid(nums, 0, len(nums)-1, (len(nums)+1)>>1)
	println(num)
	// use a simple sort to mock qs
	// sort.Slice(nums, func(i, j int) bool {
	// 	return nums[i] < nums[j]
	// })
	// num := nums[len(nums)>>1]
	for i := 0; i < len(nums); i++ {
		rsu += abs(nums[i] - num)
	}
	return rsu
}

func quickSelectMid(nums []int, low, high int, k int) int {
	// learn the quick algorithm for about 5min
	// really an awesome algorithm that have O(n) + O(n/2) + O(n/4) + ... ~= O(2n) tc
	// let's implement it:

	// choose a pivot and move it to the right
	pivot := rand.Intn(high-low) + low
	nums[pivot], nums[high] = nums[high], nums[pivot]

	// try to partition
	left := low
	for i := low; i < high; i++ {
		if nums[i] < nums[high] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[high] = nums[high], nums[left]

	if left+k-1 == high {
		return nums[left]
	}

	if left+k-1 > high {
		// bug: return quickSelectMid(nums, low, left-1, k-len(nums)+left-1)
		return quickSelectMid(nums, low, left-1, k-high+left-1)
	}

	// bug: return quickSelectMid(nums, left+1, high, k-left)
	return quickSelectMid(nums, left+1, high, k)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
