package go_test

func minCuttingCost(n int, m int, k int) int64 {
	const INF int64 = 1<<62
	ans := INF

	if n <= k && m <= k {
		return 0
	}

	if n > k && m <= k {
		x := n - k
		if x < 1 {
			x = 1
		}
		if x <= k && n-x <= k {
			cost := int64(x) * int64(n-x)
			if cost < ans {
				ans = cost
			}
		}
	}

	if m > k && n <= k {
		y := m - k
		if y < 1 {
			y = 1
		}
		if y <= k && m-y <= k {
			cost := int64(y) * int64(m-y)
			if cost < ans {
				ans = cost
			}
		}
	}

	return ans
}

