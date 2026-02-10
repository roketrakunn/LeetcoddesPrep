package go_test

import "slices"

func maxFreqSum(s string) int {

	maxV , maxC := 0 , 0

	vowels := []rune{'a','e','i','o','u'}
	freq := make(map[rune]int)

	for _ , ch := range s { 
		freq[rune(ch)]++
	}

	for k , v := range freq { 
		if slices.Contains(vowels , rune(k)) { 
			maxV = max(maxV,v)
		} else { 
			maxC = max(maxC , v)
		}
	}

	return  maxV + maxC
}
