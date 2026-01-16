package go_test

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	var b strings.Builder

	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		default:
			return false
		}
	}

	for i, w := range words {
		if i > 0 {
			b.WriteByte(' ')
		}

		if isVowel(w[0]) {
			b.WriteString(w)
		} else {
			b.WriteString(w[1:])
			b.WriteByte(w[0])
		}

		b.WriteString("ma")
		b.WriteString(strings.Repeat("a", i+1))
	}

	return b.String()
}

