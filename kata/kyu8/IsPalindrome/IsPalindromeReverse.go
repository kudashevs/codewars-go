package ispalindrome

import (
	"strings"
)

func IsPalindromeReverse(str string) bool {
	str = strings.ToLower(str)
	rev := reverseString(str)

	return str == rev
}

func reverseString(str string) (res string) {
	for _, v := range str {
		res = string(v) + res
	}

	return
}
