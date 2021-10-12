package splitstrings

import "regexp"

func SplitStringRegex(str string) []string {
	regex := regexp.MustCompile(`\w{2}`)

	return regex.FindAllString(str+`_`, -1)
}
