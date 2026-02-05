package go_test

func findLHS(nums []int) int {

	freq := make(map[int]int)

	for _ , v := range nums { 
		freq[v]++
	}
	best :=0 

	for x, c := range freq {
		if c2, ok := freq[x+1] ;ok { 
			if c + c2 > best { 
				best = c + c2
			}
		}
	}
	return best
}


