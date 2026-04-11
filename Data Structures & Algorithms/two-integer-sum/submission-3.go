func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 0)

	res:=[]int{}

	for i, v := range nums {
		
		remainder := target - v

		if valueMap, ok:= m[remainder]; ok {
			fmt.Println("halo?" , v)
			res = append(res, valueMap)
			res = append(res, i)

			sort.Ints(res)
			return res
		} else {
			m[v] = i
		}
	}

	return res
}
