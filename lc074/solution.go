package solution

func searchMatrix(matrix [][]int, target int) bool {
	var left, right, mid, anchor int
	left = 0
	right = len(matrix) - 1
	for left < right-1 {
		mid = (left + right) / 2
		if matrix[0][mid] == target {
			return true
		} else if matrix[0][mid] > target {
			right = mid
		} else {
			left = mid
		}
	}

	anchor = left
	left = 0
	right = len(matrix[0]) - 1

	for left < right-1 {
		mid = (left + right) / 2
		if matrix[mid][anchor] == target {
			return true
		} else if matrix[mid][anchor] > target {
			right = mid
		} else {
			left = mid
		}
	}

	return false
}
