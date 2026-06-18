type tempStruct struct {
	number int
	occurence int
}


func topKFrequent(nums []int, k int) []int {

	occurenceMap := make(map[int]int)

	for _, v := range nums {
		occurenceMap[v] +=1
	}

	
	tempArr := []tempStruct{}

	for k, v := range occurenceMap {
		temp := tempStruct {
			number: k,
			occurence: v,
		}

		tempArr = append(tempArr, temp)
	}

	sort.Slice(tempArr, func(i, j int) bool {
    	return tempArr[i].occurence > tempArr[j].occurence
	})


	res := []int{}
	for i:= 0; i < k; i++ {
		res = append(res, tempArr[i].number)
	}

	return res
}
