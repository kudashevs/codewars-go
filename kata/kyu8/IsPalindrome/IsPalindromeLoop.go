package ispalindrome

import (
	"strings"
)

func IsPalindromeLoop(str string) bool {
	str = strings.ToLower(str)

	first := 0
	last := len(str) - 1

	for first < last {
		if str[first] != str[last] {
			return false
		}

		first++
		last--
	}

	return true
}
