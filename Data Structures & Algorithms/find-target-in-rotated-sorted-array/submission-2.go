func binary(nums []int, l, r, target int) int {
	for l <= r {
		mid := l + (r-l) / 2

		if nums[mid] == target {
			return mid
		} 
		
		if nums[mid] < target {
			l = mid + 1
		} 

		if nums[mid] > target {
			r = mid - 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1


	// already sorted
	if nums[l] < nums[r] {
		return binary(nums, l, r, target)
	}

	// find cutting point

	for l < r {
		mid := l + (r-l) / 2

		if nums[mid] <= nums [r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	pivot := l
	fmt.Println("pivot", pivot)

	result := binary(nums, 0, pivot-1, target)
	if result != -1 {
		return result
	}

	return binary(nums, pivot, len(nums)-1, target)

}