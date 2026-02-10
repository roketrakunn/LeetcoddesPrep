package go_test

func maxScore(s string) int {
	onesRight := 0
	for _, ch := range s {
		if ch == '1' {
			onesRight++
		}
	}

	zerosLeft := 0
	best := -1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zerosLeft++
		} else { 
			onesRight--
		}
		if zerosLeft+onesRight > best {
			best = zerosLeft + onesRight
		}
	}
	return best
}



