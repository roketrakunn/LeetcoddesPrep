package go_test

func earliestTime(tasks [][]int) int {

	minFinish := tasks[0][0] + tasks[0][1]

	for _, task := range tasks {
		finish := task[0] + task[1]
		if finish < minFinish {
			minFinish = finish
		}
	}
	return minFinish
}

