package solution

func FindPoisonedDuration(timeSeries []int, duration int) int {
	rsu := 0
	currentBegin := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] <= duration {
			continue
		}
		rsu += timeSeries[i-1] - timeSeries[currentBegin] + duration
		currentBegin = i
	}
	rsu += timeSeries[len(timeSeries)-1] - timeSeries[currentBegin] + duration

	return rsu
}
