package go_test


func validPalindrome(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1

	for i < j {
		if r[i] == r[j] {
			i++
			j--
			continue
		}
		return isPalRange(r, i+1, j) || isPalRange(r, i, j-1)
	}
	return true
}

func isPalRange(r []rune, i, j int) bool {
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}

