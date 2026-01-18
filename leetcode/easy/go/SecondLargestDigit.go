package go_test

func secondHighest(s string) int {
	first,second := -1 , -1

	for  _,cha := range s { 
		
		if cha >= '0' && cha <= '9' { 
			d := int(cha -'0')

			if d > first { 
				second = first
				first = d
			} else if d < first && d > second {
				second = d
			}
		}
	}
	return second
}

