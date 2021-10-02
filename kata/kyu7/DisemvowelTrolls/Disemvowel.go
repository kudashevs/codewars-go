package disemvoweltrolls

import "regexp"

func Disemvowel(str string) string {
	regex := regexp.MustCompile(`(?i)[aeiou]`)

	return regex.ReplaceAllString(str, "")
}
