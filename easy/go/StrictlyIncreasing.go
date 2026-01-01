
package mostFrequent

func canBeIncreasing(nums []int) bool {

	allowed := 0 

	for i ,val := range nums { 
		for allowed < 2 { 
			for j := i+1 ; j < len(nums); j++{ 
				if val >= nums[j] { 
					allowed++
				}
			}
		}
	}
	return  allowed == 1 || allowed == 0
}
