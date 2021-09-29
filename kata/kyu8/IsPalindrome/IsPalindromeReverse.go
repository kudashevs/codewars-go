package ispalindrome

import (
	"strings"
)

func IsPalindromeReverse(str string) bool {
	str = strings.ToLower(str)
	rev := reverse(str)

	return str == rev
}

func reverse(str string) (res string) {
	for _, v := range str {
		res = string(v) + res
	}

	return
}
