func maxProfit(prices []int) int {


	left := 0
	right := left + 1

	result := 0

	for right < len(prices){
		profit := prices[right] - prices[left]

		if profit > result {
			result = profit
		}

		if prices[left] > prices[right] {
			left = right
			right = left + 1
		} else {
			right+=1
		}
	}

	return result
}
