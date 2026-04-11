func maxArea(heights []int) int {

	left := 0
	right := len(heights) - 1
	result := 0

	for {
		if left >= right {
			break
		}

		// calculate current area

		height := heights[left]

		if heights[right] < height {
			height = heights[right]
		}

		area := height * (right - left)

		if result < area {
			result = area
		}

		if heights[left] <= heights[right] {
			left+=1
		} else {
			right-=1
		}
	}

	return result
}
