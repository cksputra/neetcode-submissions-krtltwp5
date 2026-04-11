func isPalindrome(s string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]+") 
	cleanedString := re.ReplaceAllString(s, "")
	cleanedString = strings.ToLower(cleanedString)

	palindrome := true
	for i:=0; i < len(cleanedString)/2; i++ {
		lastNum := len(cleanedString)-1-i

		if cleanedString[i] != cleanedString[lastNum] {

			left := cleanedString[i]
			right := cleanedString[lastNum]

			fmt.Println(left, right)
			palindrome = false
			break
		} 
	}

	return palindrome
}

