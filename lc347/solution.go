package solution

type node struct {
	val, cnt int
}

func topKFrequent(nums []int, k int) []int {
	rsu := []int{}
	heap := []*node{}
	numsPtr := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if targetIdx, ok := numsPtr[nums[i]]; ok {
			heap[targetIdx].cnt++
			for targetIdx >= 1 && heap[targetIdx].cnt > heap[(targetIdx-1)>>1].cnt {
				numsPtr[heap[(targetIdx-1)>>1].val] = targetIdx
				heap[targetIdx], heap[(targetIdx-1)>>1] = heap[(targetIdx-1)>>1], heap[targetIdx]
				targetIdx = (targetIdx - 1) >> 1
			}
			numsPtr[nums[i]] = targetIdx
		} else {
			heap = append(heap, &node{val: nums[i], cnt: 1})
			numsPtr[nums[i]] = len(heap) - 1
		}
	}

	for i := 0; i < k; i++ {
		rsu = append(rsu, heap[i].val)
	}
	return rsu
}
