package go_test

func isPossibleToSplit(nums []int) bool {
	freq := make(map[int]int)

	for _, x := range nums {
		freq[x]++
		if freq[x] > 2 {
			return false
		}
	}
	return true
}

