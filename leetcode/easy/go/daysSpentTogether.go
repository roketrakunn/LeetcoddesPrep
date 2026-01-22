package go_test

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	parts := strings.Split(date, "-")
	month, _ := strconv.Atoi(parts[0])
	day, _ := strconv.Atoi(parts[1])

	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	total := 0
	for i := 0; i < month-1; i++ {
		total += daysInMonth[i]
	}

	return total + day
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aStart := dayOfYear(arriveAlice)
	aEnd := dayOfYear(leaveAlice)
	bStart := dayOfYear(arriveBob)
	bEnd := dayOfYear(leaveBob)

	start := max(aStart, bStart)
	end := min(aEnd, bEnd)

	if start > end {
		return 0
	}
	return end - start + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

