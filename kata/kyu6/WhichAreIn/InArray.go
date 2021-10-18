package whicharein

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	result := []string{}
	tmp := map[string]bool{}

	for _, v := range array1 {
		if findInArray(v, array2) {
			tmp[v] = true
		}
	}

	for key := range tmp {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}

func findInArray(str string, array []string) bool {
	for _, v := range array {
		if strings.Contains(v, str) {
			return true
		}
	}

	return false
}
