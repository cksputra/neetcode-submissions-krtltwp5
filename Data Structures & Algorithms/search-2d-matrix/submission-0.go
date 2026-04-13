func searchMatrix(matrix [][]int, target int) bool {
	extend := []int{}
	for i := range matrix {
		extend = append(extend, matrix[i]...)
	}

	left := 0
	right := len(extend) - 1

	for left <= right {
		mid := left + (right-left) / 2

		if extend[mid] == target {
			return true
		}

		if extend[mid] > target {
			right = mid - 1
		}

		if extend[mid] < target {
			left = mid + 1
		}
	}

	return false
}
