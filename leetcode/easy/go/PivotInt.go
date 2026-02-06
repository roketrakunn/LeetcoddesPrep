package go_test

import "math"

func pivotInteger(n int) int {
	total := n * (n + 1) / 2
	root := int(math.Sqrt(float64(total)))
	if root*root == total {
		return root
	}
	return -1
}

