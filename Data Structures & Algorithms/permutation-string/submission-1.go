func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	alphabetArr, s2Arr := [26]int{}, [26]int{}

	for _, v := range s1 {
		alphabetArr[v - 'a']+=1
	}

	l:=0

	for r := 0; r < len(s2);r++ {
		// fmt.Println(alphabetArr)
		// fmt.Println(s2Arr)
		// fmt.Println("====")

		s2Arr[s2[r] - 'a']+=1

		if r > len(s1) - 1 {
			s2Arr[s2[l] - 'a']-=1
			l++
		}

		if alphabetArr == s2Arr {
			return true
		}
	}


	
	return false
}

