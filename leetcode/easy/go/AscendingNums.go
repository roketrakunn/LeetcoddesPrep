package go_test

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	tokens := strings.Split(s, " ")

	prev := -1

	for _, t := range tokens {
		if num, err := strconv.Atoi(t); err == nil {
			if num <= prev {
				return false
			}
			prev = num
		}
	}

	return true
}

