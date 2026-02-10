package go_test

func removeZeros(n int64) int64 {
	var res int64 = 0
	var mul int64 = 1

	for n > 0 {
		d := n % 10
		if d != 0 {
			res += d * mul
			mul *= 10
		}
		n /= 10
	}

	return res
}

