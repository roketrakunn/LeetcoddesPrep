package go_test


func leftRightDifference(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	total := 0
	for _, v := range nums {
		total += v
	}

	left := 0
	right := total

	for i, v := range nums {
		right -= v
		diff := left - right
		if diff < 0 {
			diff = -diff
		}
		res[i] = diff
		left += v
	}

	return res
}

