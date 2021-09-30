package vowelcount

import "regexp"

func GetCountRegex(str string) (count int) {
	regex := regexp.MustCompile(`[aeiou]`)

	return len(regex.FindAllString(str, -1))
}
