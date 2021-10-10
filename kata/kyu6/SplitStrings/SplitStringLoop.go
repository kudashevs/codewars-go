package splitstrings

import "strings"

func SplitStringLoop(str string) []string {
	res := []string{}

	if len(str) == 0 {
		return res
	}

	chars := strings.Split(str, ``)
	part := ``

	for i, v := range chars {
		part += v

		if i%2 != 0 {
			res = append(res, part)
			part = ``
		}
	}

	if len(part) > 0 {
		res = append(res, part+`_`)
	}

	return res
}
