func groupAnagrams(strs []string) [][]string {

	mapLength := make(map[int][]string,0)

	runeMap := make(map[string]map[rune]int, 0)

	doneMap := make(map[string]bool, 0)

	res := [][]string{}

	// preprocessing
	for _, str := range strs {
		length := len(str)
		mapLength[length] = append(mapLength[length], str)

		tempMap := make(map[rune]int, 0)
		for _, char := range str {
			tempMap[char]+=1
		}

		runeMap[str] = tempMap
	}	

	// grouping
	for _, arrStr := range mapLength {
		for i:=0; i<len(arrStr);i++{
			// skip if word already have a group
			if doneMap[arrStr[i]] {
				continue
			}

			// init array with current word
			resultArr := []string{arrStr[i]}

			for j:=i+1; j<len(arrStr);j++{
				anagram := true
				for k, _ := range runeMap[arrStr[i]] {
					// compare map
					if runeMap[arrStr[i]][k] != runeMap[arrStr[j]][k] {
						anagram = false
						break
					}
				}

				if anagram {
					doneMap[arrStr[j]] = true
					resultArr = append(resultArr, arrStr[j])
				}
			}
			
			res = append(res, resultArr)
		}
	}

	fmt.Println(res)

	return res
}
