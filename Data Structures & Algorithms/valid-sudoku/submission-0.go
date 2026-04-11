func getBoxPlacement(x, y int) int {
	if x < 3 {
		if y < 3 {
			return 0
		} else if y < 6 {
			return 1
		} else {
			return 2
		}
	}

	if x >= 3 && x < 6 {
		if y < 3 {
			return 3
		} else if y < 6 {
			return 4
		} else {
			return 5
		}
	}

	if x >= 6 && x < 9 {
		if y < 3 {
			return 6
		} else if y < 6 {
			return 7
		} else {
			return 8
		}
	}
	return 0
}

func isValidSudoku(board [][]byte) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // Handles the panic
		}
	}()


	rowMap := make([]map[string]bool, 9)
	for i := range rowMap {
		rowMap[i] = make(map[string]bool)
	}
	colMap := make([]map[string]bool, 9)
	for i := range colMap {
		colMap[i] = make(map[string]bool)
	}
	boxMap := make([]map[string]bool, 9) // left to right, 0 1 2, 3 4 5, ...
	for i := range boxMap {
		boxMap[i] = make(map[string]bool)
	}
	valid := true

	for x := range board {
		for y := range board[x] {
			num := string(board[x][y])
			if num == "." {
				continue
			}
			// check num exist in row or not
			if rowMap[x][num] == true {
				fmt.Println("row break", x, y)
				valid = false
				break
			} else {
				rowMap[x][num] = true
			}

			// check if num already exist in column
			if colMap[y][num] == true {
				valid = false
				fmt.Println("column break", x, y)
				break
			} else {
				colMap[y][num] = true
			}

			// check if num already exist in box
			index := getBoxPlacement(x,y)

			if boxMap[index][num] == true {
				fmt.Println("box break", x, y)
				valid = false
				break
			} else {
				boxMap[index][num] = true
			}
		}

		if !valid {
			break
		}
	}

	return valid
}



