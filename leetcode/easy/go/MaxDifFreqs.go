package go_test

func maxDifference(s string) int {
	freqs := make(map[rune]int)

	for _, ch := range s {
		freqs[ch]++
	}

	max := 0
	min := -1

	for _, v := range freqs {
		if v > max && v % 2 != 0 {
			max = v
		}
		if (min == -1 || v < min)  && v % 2 == 0 {
			min = v
		}
	}

	return max - min
}

