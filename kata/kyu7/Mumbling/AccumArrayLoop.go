package mumbling

import "strings"

func AccumArrayLoop(s string) string {
	chars := strings.Split(s, ``)
	result := make([]string, len(chars))

	for i, v := range chars {
		ch := strings.ToLower(v)
		result[i] = strings.ToUpper(ch) + strings.Repeat(ch, i)
	}

	return strings.Join(result, `-`)
}
