package go_test

func countSubarrays(nums []int) int {
	count := 0

	for i := 0; i+2 < len(nums); i++ {
		if 2*(nums[i]+nums[i+2]) == nums[i+1] {
			count++
		}
	}

	return count
}
