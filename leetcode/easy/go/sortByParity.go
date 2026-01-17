package go_test

func sortArrayByParity(nums []int) []int {
	sorted := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for _, val := range nums {
		if val%2 == 0 {
			sorted[left] = val
			left++
		} else {
			sorted[right] = val
			right--
		}
	}

	return sorted
}

