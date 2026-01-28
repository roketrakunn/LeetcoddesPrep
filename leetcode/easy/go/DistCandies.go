package go_test

func distributeCandies(candyType []int) int{ 
	
	seen := make(map[int]bool)
	count := 0

	for _,val := range candyType { 
		if !seen[val] { 
			seen[val] = true
			count ++
			if count == len(candyType) / 2{ 
				break
			}
		}
	}
	return count
}

