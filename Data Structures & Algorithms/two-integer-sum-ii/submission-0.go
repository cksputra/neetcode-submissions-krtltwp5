func twoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1

	for {
		if numbers[left] + numbers[right] == target {
			return []int{left+1, right+1}
		}

		if numbers[left] + numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	
	return []int{}
}
