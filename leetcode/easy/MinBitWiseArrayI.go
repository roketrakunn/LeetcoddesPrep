package go_test

func minBitwiseArray(nums []int) []int {
	res := make([]int, len(nums))

	for i, x := range nums {
		if x%2 == 0 {
			res[i] = -1
			continue
		}

		t := 0
		tmp := x
		for tmp&1 == 1 {
			t++
			tmp >>= 1
		}
		res[i] = x - (1 << (t - 1))
	}

	return res
}

