func isValid(s string) bool {
	stack := []string{}

	valid := false

	bracketMap := map[string]string{
		"]" : "[",
		")" : "(",
		"}" : "{",
	}

	for _, v := range s {
		str := string(v)
		if str == "]" || str == "}" || str == ")" {
			if len(stack) == 0 {
				return false
			}
			
			// compare stack
			if stack[len(stack)-1] != bracketMap[str] {
				return false
			}

			// pop stack
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str)
		}
	}

	if len(stack) == 0 {
		valid = true
	} 

	return valid
}
