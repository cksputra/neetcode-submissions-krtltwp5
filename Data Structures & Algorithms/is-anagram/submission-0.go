func isAnagram(s string, t string) bool {
	m1 := make(map[rune]int)
	
	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		m1[v]+=1
	}

	for _, v := range t {
		m1[v]-=1
	}

	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}

	return true
}
