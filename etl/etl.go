package etl

import "strings"

func Transform(data map[int][]string) (result map[string]int) {
	result = make(map[string]int)
	for k, vals := range data {
		for _, v := range vals {
			new_key := strings.ToLower(v)
			result[new_key] += k
		}
	}
	return
}
