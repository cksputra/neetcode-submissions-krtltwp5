type attr struct {
	Index int
	Value int
}

func dailyTemperatures(temperatures []int) []int {


	length := len(temperatures)
	
	res := make([]int, length)

	stack := []attr{}

	for i, v := range temperatures {
		temp := attr {
			Index: i,
			Value: v,
		}

		if len(stack) == 0 {
			stack = append(stack, temp)
		} else {
			for {
				if len(stack) == 0 {
					break
				}

				//get top
				top := stack[len(stack) - 1]

				// stop when top is bigger than current
				if top.Value >= v {
					break
				}

				//insert into array
				res[top.Index] = i - top.Index

				//pop
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, temp)
		}
	}

	return res
}

