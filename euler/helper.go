package euler

func reverse(str string) string {
	length := len(str)
	runes := []rune(str)
	for i := 0; i < length/2; i++ {
		runes[length-1-i], runes[i] = runes[i], runes[length-1-i]
	}
	return string(runes)
}

func isPalindrome(str string) bool {
	result := str == reverse(str)
	return result
}
