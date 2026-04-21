package go_test

func canAliceWin(nums []int) bool {
    
	sum_single := 0 ;
	sum_doubles :=0 ; 

	for _, val := range nums { 
		if val >= 0 &&  val <= 9 { 
			sum_single += val
		}else { 
			sum_doubles +=val
		}
	}

	if sum_doubles > sum_single  || sum_single > sum_doubles{ 
		return true
	}
	return false
}
