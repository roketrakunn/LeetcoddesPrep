package go_test

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}


