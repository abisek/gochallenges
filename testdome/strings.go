package testdome

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	runeCount := make(map[rune]int)
	for _, runeValue := range str1 {
		if count, ok := runeCount[runeValue]; ok {
			runeCount[runeValue] = count + 1
		} else {
			runeCount[runeValue] = 0
		}
	}

	for _, runeValue := range str2 {
		if count, ok := runeCount[runeValue]; ok {
			if count == 1 {
				delete(runeCount, runeValue)
			} else {
				runeCount[runeValue] = count - 1
			}
		} else {
			return false
		}
	}

	return true
}
