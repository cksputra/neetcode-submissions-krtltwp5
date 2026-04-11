func productExceptSelf(nums []int) []int {

	res := []int{}
	prefix := []int{}
	suffix := []int{}

	pre := nums[0]
	for i, v := range nums {
		// skip first element
		if i == 0 {
			prefix = append(prefix, 0)
			continue
		}
		prefix = append(prefix, pre)
		pre = pre * v
	}

	suffix = append(suffix, 0)
	suf := nums[len(nums)-1]
	// start from -2 because last element is always 0
	for i:=len(nums) - 2;  i>=0; i--{
		suffix = append([]int{suf}, suffix...)
		suf = suf * nums[i]
	}

	fmt.Println(prefix)
	fmt.Println(suffix)

	for i := range nums {
		a := prefix[i]
		b := suffix[i]

		if i == 0 {
			a = 1
		}

		if i == len(nums) - 1 {
			b = 1
		}

		res = append(res, a * b)
	}

	return res
}