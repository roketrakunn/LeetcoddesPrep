package go_test

import (
	"slices"
)

func minMovesToSeat(seats []int, students []int) int {
	count := 0

	slices.Sort(seats) // this is O(n)
	slices.Sort(students) //and so is this 

	for i , seat := range seats { 
		count +=  max(seat, students[i]) - min(seat, students[i])
	}
	return count
}
