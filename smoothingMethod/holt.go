package smoothingMethod

func getHoltTrend(data []float64, T int, alpha, beta float64) float64 {
	var B float64
	if T == 1 {
		B = data[1] - data[0]
		return B
	}
	B = (beta * (getHoltLevel(data, T, alpha, beta) - getHoltLevel(data, T-1, alpha, beta))) + ((1 - beta) * getHoltTrend(data, T-1, alpha, beta))

	return B
}

func getHoltLevel(data []float64, T int, alpha, beta float64) float64 {
	var level float64
	if T == 1 {
		level = data[0]
		return level
	}

	level = (alpha * data[T]) + (1-alpha)*(getHoltLevel(data, T-1, alpha, beta)+getHoltTrend(data, T-1, alpha, beta))

	return level
}

func HoltPredict(level, trend float64, after int) float64 {
	return 0.0
}
