func AbsInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func findDuplicate(nums []int) int {
    for _, v := range nums {
		index := AbsInt(v) - 1

		if nums[index] < 0 {
			return AbsInt(v)
		}

		nums[index] = nums[index] * -1
	}

	return 0
}
