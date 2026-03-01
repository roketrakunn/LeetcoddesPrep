package go_test

func isPrefixString(s string, words []string) bool {
	curr := ""
	k := len(s)

	for _, word := range words {
		if len(curr) >= k {
			return false
		}
		curr += word
		if curr == s {
			return true
		}
	}
	return false
}
