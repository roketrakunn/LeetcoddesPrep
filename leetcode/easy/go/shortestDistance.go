package go_test
import "math"

func shortestToChar(s string, c byte) []int {
	n := len(s)
	result := make([]int, n)

	prev := -math.MaxInt32

	for i := 0; i < n; i++ {
		if s[i] == c {
			prev = i
		}
		result[i] = i - prev
	}

	prev = math.MaxInt32

	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		result[i] = min(result[i], prev - i)
	}

	return result
}
