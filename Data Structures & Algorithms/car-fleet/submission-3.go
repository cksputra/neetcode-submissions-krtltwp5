import (
	"slices"
)

type car struct {
	Position int
	Speed int
}

func carFleet(target int, position []int, speed []int) int {

	carArr := []car{}
	for i := range position {
		carArr = append(carArr, car{
			Position: position[i],
			Speed: speed[i],
		})
	}
	slices.SortFunc(carArr, func(a, b car) int {
    	return b.Position - a.Position // Descending order
	})

	timeArr := []float64{}

	for _, v := range carArr {
		time := float64(target-v.Position)/ float64(v.Speed)

		timeArr = append(timeArr, time)
	}

	stack := []float64{}

	
	for _, v := range timeArr {
		curr := v
		if len(stack) == 0 {
			stack = append(stack, curr)
		} else {
			now := stack[len(stack) - 1]

			if curr != now && now < curr{
				stack = append(stack, curr)
			}
		}
	}

	return len(stack)
}

// 10 8 5 3 0
// 2  4 1 3 1
// 1. 1 7 3 11

// target = 12


// 8 7 6 5 4 3

// 10

// 2 / 4 = 0.5
// 3 / 4 = 
