package smoothingMethod

func DoubleMovingAverage(ma []float64, N int, T int) float64 {
	var sum float64
	var dma float64

	for i := T - N + 1; i <= T; i++ {
		if T-N+1 < N-1 {
			return 0.0
		}
		if ma[i] == 0 {
			continue
		}
		sum += ma[i]
	}
	dma = sum / float64(N)

	return dma
}

func DoubleMovingAverages(ma []float64, N int) []float64 {
	var result []float64
	for i := 0; i < len(ma); i++ {
		temp := DoubleMovingAverage(ma, N, i)

		result = append(result, temp)
	}
	return result
}
