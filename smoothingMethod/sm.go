package smoothingMethod

import (
	"fmt"
	"math"

	"github.com/evaluation"
)

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

func EWMA(data []float64, alpha float64, T int) float64 {
	if T == 0 {
		return data[0]
	}
	result := (alpha * data[T]) + ((1 - alpha) * EWMA(data, alpha, T-1))
	return result
}

func EWMAs(data []float64, alpha float64) []float64 {
	var ewma []float64
	for i := 0; i < len(data); i++ {
		ewma = append(ewma, EWMA(data, alpha, i))
	}
	return ewma
}

func GetAlpha(training, test []float64, T int) (alpha float64) {

	var min float64 = math.MaxFloat64
	findAlpha := make(map[float64]float64)
	predicted := make([]float64, len(test))
	for alpha = 0.01; alpha < 1; alpha += 0.01 {
		ewma := EWMA(training, alpha, T)
		for i := 0; i < len(test); i++ {
			predicted[i] = ewma
		}
		temp := evaluation.MSE(test, predicted)
		if min > temp {
			min = temp
		}
		findAlpha[temp] = alpha
	}

	return findAlpha[min]
}

func GetAlphaOfBrown(training, test []float64, T int) (alpha float64) {

	var min float64 = math.MaxFloat64
	findAlpha := make(map[float64]float64)
	predicted := make([]float64, len(test))
	for alpha = 0.01; alpha < 1; alpha += 0.01 {
		ewma := DoubleExponentialSmoothing(training, alpha, T)
		for i := 0; i < len(test); i++ {
			predicted[i] = ewma
		}
		temp := evaluation.MSE(test, predicted)
		if min > temp {
			min = temp
		}
		findAlpha[temp] = alpha
	}

	return findAlpha[min]
}

//Xt = c + Bt + at

//선형추세와 이중지수평활

func DoubleExponentialSmoothing(ewma []float64, alpha float64, T int) float64 {
	if T == 0 {
		return ewma[0]
	}
	return (alpha * ewma[T]) + (float64(1)-alpha)*DoubleExponentialSmoothing(ewma, alpha, T-1)
}

func Brown(ewma []float64, alpha float64) []float64 {
	var brown []float64
	for i := 0; i < len(ewma); i++ {
		brown = append(brown, DoubleExponentialSmoothing(ewma, alpha, i))
	}
	return brown
}

func GetTrendOfBrown(ewma, brown []float64, alpha float64, T int) (B float64) {
	B = (alpha / (1 - alpha)) * (ewma[T] - brown[T])

	return B
} //맞음

func GetConstantOfBrown(ewma, brown []float64, alpha float64, T int) (C float64) {
	C = (2 * ewma[T]) - brown[T] - (GetTrendOfBrown(ewma, brown, alpha, T) * float64(T))
	return C
} //틀림

func PrintFormulaOfBrown(ewma, brown []float64, alpha float64, T int) {
	fmt.Printf("Xt = %f + %f * t + at (at는 오차항)",
		GetConstantOfBrown(ewma, brown, alpha, T), GetTrendOfBrown(ewma, brown, alpha, T))
}
