package mumbling

import "strings"

func AccumStringLoop(s string) string {
	result := strings.ToUpper((s[:1]))

	for i := 1; i < len(s); i++ {
		ch := strings.ToLower(string(s[i]))
		result += `-` + strings.ToUpper(ch) + strings.Repeat(ch, i)
	}

	return result
}
