func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)

	l := 0

	currLongest := 0
	res := 0
	for r:=0; r < len(s); r++ {
		freqMap[s[r]]+=1

		if freqMap[s[r]] > currLongest {
			currLongest = freqMap[s[r]]
		}

		for ((r-l+1) - currLongest > k) {
			freqMap[s[l]]--
			l++
		}

		window := r - l + 1
		if  window > res {
			res = window
		}
	}

	return res
}


