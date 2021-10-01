package vowelcount

import "strings"

func GetCountLoop(str string) (count int) {
	vowels := `aeiou`

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return
}
