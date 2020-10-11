package solution

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var left, right, mid, prev int
	left = 0
	right = len(matrix)
	for left < right-1 {
		mid = left + (right-left)>>1
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			right = mid
		} else {
			left = mid
		}
	}

	prev = left
	right = len(matrix[0]) - 1
	left = 0
	for left < right-1 {
		mid = left + (right-left)>>1
		if matrix[prev][mid] == target {
			return true
		} else if target > matrix[prev][mid] {
			left = mid
		} else {
			right = mid
		}
	}

	right = len(matrix[0]) - 1
	left = 0
	for left < right-1 {
		mid = left + (right-left)>>1
		if matrix[0][mid] == target {
			return true
		} else if matrix[0][mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	prev = left
	left = 0
	right = len(matrix) - 1
	for left < right-1 {
		mid = left + (right-left)>>1
		if matrix[mid][prev] == target {
			return true
		} else if target > matrix[mid][prev] {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}
