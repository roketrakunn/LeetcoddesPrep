package go_test

import (
	"slices"
)

func minMovesToSeat(seats []int, students []int) int {
	count := 0

	slices.Sort(seats)
	slices.Sort(students)

	for i , seat := range seats { 
		count +=  max(seat, students[i]) - min(seat, students[i])
	}
	return count
}
