package go_test

func maximumStrongPairXor(nums []int) int {
	maxXor := 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			x, y := nums[i], nums[j]
			if absInt(x-y) <= minInt(x, y) {
				if v := x ^ y; v > maxXor {
					maxXor = v
				}
			}
		}
	}

	return maxXor
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

