package go_test

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	sep := string(separator)
	res := []string{}

	for _, str := range words {
		parts := strings.Split(str, sep)
		for _, p := range parts {
			if p != "" {
				res = append(res, p)
			}
		}
	}
	return res
}

