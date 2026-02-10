package go_test

func stableMountains(height []int, threshold int) []int {

	res := []int{}

	for i := 1 ; i < len(height) ; i++ { 
		prev := height[i-1]
		if prev > threshold { 
			res = append(res, i)
		} 
	} 
	return  res
}
