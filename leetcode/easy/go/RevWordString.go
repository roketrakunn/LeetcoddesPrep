package go_test

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i , w := range words { 
		words[i] = reverse(w)
	}

	return  strings.Join(words, " ")
}

func reverse(s string) string { 
	r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
