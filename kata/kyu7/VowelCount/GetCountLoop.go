package vowelcount

import "strings"

func GetCountLoop(str string) (count int) {
	vowelsCount := 0
	vowels := `aeiou`

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			vowelsCount++
		}
	}

	return vowelsCount
}
