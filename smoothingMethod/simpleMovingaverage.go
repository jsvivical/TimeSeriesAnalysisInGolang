package smoothingMethod

func movingAverage(data []float64, N, T int) float64 {
	// n은 이동평균을 구할 때 사용되는 데이터의 개수
	//time은 구하려는 시점(인덱스)
	var sum float64
	var mv float64

	for i := T - N + 1; i <= T; i++ {
		if T-N+1 < 0 {
			return 0.0
		}
		sum += data[i]
	}
	mv = sum / float64(N)

	return mv
}

func MovingAverages(data []float64, N int) []float64 {
	var result []float64
	for i := 0; i < len(data); i++ {
		temp := movingAverage(data, N, i)

		result = append(result, temp)
	}
	return result
}

func PredByMA(MA []float64, t, N int) float64 {

	if t-N < 0 {
		return 0
	}

	pred := MA[t] + (MA[t]-MA[t-N])/float64(N)
	return pred
}
