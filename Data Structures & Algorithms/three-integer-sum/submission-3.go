func threeSum(nums []int) [][]int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // Handle the panic value
		}
	}()


	result := make([][]int, 0)

	hash := make(map[[3]int]bool, 0)

	sort.Ints(nums)

	for i, num := range nums {

		j := i+1
		if j == len(nums) {
			break
		}

		k := len(nums) - 1

		for j != k {
			if num + nums[j] + nums[k] == 0 {
				temp := [3]int{nums[i], nums[j], nums[k]}
				if hash[temp] {
					j++
				} else {
					hash[temp] = true
					result = append(result, temp[:])
					j++
				}
			}else if num + nums[j] + nums[k] > 0 {
				k--
			} else {
				j++
			}
		} 
	}

	return result
}

// -4, -1,-1,0, 1,2

// -4, -3, -2, -1, -1, 0, 0, 1, 2, 3, 4