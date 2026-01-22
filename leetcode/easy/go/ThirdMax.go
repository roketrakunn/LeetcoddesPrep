package go_test

func thirdMax(nums []int) int {
	var first, second, third *int

	for _, v := range nums {
		// skip duplicates
		if (first != nil && v == *first) ||
			(second != nil && v == *second) ||
			(third != nil && v == *third) {
			continue
		}

		if first == nil || v > *first {
			third = second
			second = first
			first = &v
		} else if second == nil || v > *second {
			third = second
			second = &v
		} else if third == nil || v > *third {
			third = &v
		}
	}

	if third == nil {
		return *first
	}
	return *third
}

