package utils

import "sort"

// call func f ordered by keys of map m
func SortAndCall(m map[string]string, f func(k, v string)) {
	// sort keys
	keys := []string{}
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// call in key order
	for _, key := range keys {
		f(key, m[key])
	}
}
