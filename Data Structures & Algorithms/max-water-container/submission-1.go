func min (a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(heights []int) int {

	l := 0
	r := len(heights) - 1

	res := 0

	for {
		if l == r {
			break
		}

		length := r - l

		height := min(heights[l], heights[r])

		sum := height * length

		if sum > res {
			res = sum
		}

		if heights[l] <= heights[r] {
			l+=1
		} else {
			r-=1
		}
	}

	return res

}
