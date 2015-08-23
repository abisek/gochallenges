package euler

import "math"

// Returns a reversed string
func Reverse(str string) string {
	length := len(str)
	resultRunes := make([]rune, length)
	for i, runeValue := range str {
		resultRunes[length-1-i] = runeValue
	}

	return string(resultRunes)
}

func IsPalindrome(str string) bool {
	result := str == Reverse(str)
	return result
}

func IsPrime(num int) bool {
	if num < 0 {
		panic("num should be non-negative")
	} else if num == 1 {
		return false
	} else if num == 2 {
		return true
	}
	upperLimit := int(math.Ceil(math.Sqrt(float64(num))))

	for i := 2; i <= upperLimit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
