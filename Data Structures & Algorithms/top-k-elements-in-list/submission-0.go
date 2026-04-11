type temp struct {
	Number int
	Occurence int
}

func topKFrequent(nums []int, k int) []int {
	
	mapNum := make(map[int]int, 0)
	for _, num := range nums {
		mapNum[num]+=1
	}

	arrTemp := []temp{}

	for k, v := range mapNum {
		tempstruct := temp{
			Number: k,
			Occurence: v,
		}

		arrTemp = append(arrTemp, tempstruct)
	}

	sort.Slice(arrTemp, func(i, j int) bool {
    	return arrTemp[i].Occurence > arrTemp[j].Occurence
	})

	// fmt.Println(arrTemp)

	res := []int{}

	for i:=0; i < k;i++{
		res= append(res, arrTemp[i].Number)
	}

	return res
}
