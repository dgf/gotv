package utils

import "sort"

// SortAndCall func f by keys of map
func SortAndCall(m map[string]string, f func(k, v string)) {
	// sort keys
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// call in key order
	for _, key := range keys {
		f(key, m[key])
	}
}
