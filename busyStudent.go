package busy_student

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for idx := range startTime {
		if startTime[idx] <= queryTime && endTime[idx] >= queryTime {
			count++
		}
	}
	return count
}
