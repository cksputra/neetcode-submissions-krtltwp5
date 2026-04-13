func minEatingSpeed(piles []int, h int) int {
	max := 0

	for _, v := range piles {
		if v > max {
			max = v
		}
	}

	left := 1
	right := max
	res := 0

	for left <= right {
		mid := left + (right-left) / 2

		time := 0
		for _, v := range piles {
			time = time + int(math.Ceil(float64(v)/ float64(mid)))
		}

		if time <= h {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
