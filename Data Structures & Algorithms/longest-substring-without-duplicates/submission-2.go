func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	left := 0
	right := left + 1

	hashMap := make(map[byte]bool)

	hashMap[s[0]] = true

	result :=1

	for left < len(s){
		// fmt.Println(left,right)
		if right == left {
			right+=1
		}

		if right >= len(s) {
			break
		}

		if !hashMap[s[right]]{
			length := right - left + 1
			if result < length {
				result = length
			}
		} else {
			for hashMap[s[right]] {
				hashMap[s[left]] = false
				left+=1
			}			
			
		}
		
		hashMap[s[right]] = true
		right+=1
	}

	return result
}


// zxyx
// xyx
// yx