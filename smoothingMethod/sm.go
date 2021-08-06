package smoothingMethod

import "fmt"

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

//이중이동평균법
//X = c + bt + a -> b,c는 상수, a는 오차항

//b먼저 구해야함 -> 트렌드를 나타내는 상수

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

func GetTrend(ma, dma []float64, N, T int) float64 {
	var B float64
	B = 2 * (float64(1) / float64(N-1)) * (ma[T] - dma[T])

	return B

}

func GetConstant(ma, dma []float64, B float64, T int) float64 {
	var C float64
	C = (2 * ma[T]) - (dma[T]) - (B * float64(T))

	return C
}

func PrintFormula(ma, dma []float64, N, T int) {
	var C, B float64
	B = GetTrend(ma, dma, N, T)
	C = GetConstant(ma, dma, B, T)

	fmt.Printf("Xt = %.4f + %.4ft + at\n\n", C, B)

}

func Predict(ma, dma []float64, N, T, after int) []float64 {
	var predVal []float64
	B := GetTrend(ma, dma, N, T)
	C := GetConstant(ma, dma, B, T)

	for i := T; i <= T+after; i++ {
		temp := C + (B * float64(i))
		predVal = append(predVal, temp)
	}

	return predVal
}
