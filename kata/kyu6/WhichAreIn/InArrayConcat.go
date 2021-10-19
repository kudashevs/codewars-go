package whicharein

import (
	"sort"
	"strings"
)

func InArrayConcat(array1 []string, array2 []string) []string {
	result := []string{}
	tmp := map[string]bool{}
	hay := strings.Join(array2, ` `)

	for _, needle := range array1 {
		if strings.Contains(hay, needle) {
			tmp[needle] = true
		}
	}

	for key := range tmp {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}
