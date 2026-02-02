package go_test

func findGCD(nums []int) int {
	minV, maxV := Min(nums), Max(nums)
	return gcd(minV, maxV)
}

func Min(arr []int) int {
	minV := arr[0]
	for _, v := range arr[1:] {
		if v < minV {
			minV = v
		}
	}
	return minV
}

func Max(arr []int) int {
	maxV := arr[0]
	for _, v := range arr[1:] {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

