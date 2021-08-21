package smoothingMethod

import (
	"fmt"
)

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
	C = (2 * ewma[T]) - brown[T] - (GetTrendOfBrown(ewma, brown, alpha, T) * float64(T+1))
	return C
} //틀림

func PrintFormulaOfBrown(ewma, brown []float64, alpha float64, T int) {
	fmt.Printf("Xt = %f + %f * t + at (at는 오차항)",
		GetConstantOfBrown(ewma, brown, alpha, T), GetTrendOfBrown(ewma, brown, alpha, T))
}

func BrownPredict(ewma, brown []float64, alpha float64, year []int, T, after int) []float64 {
	var predict []float64
	fmt.Printf("%d년의 데이터로 계산한 예측식 : ", year[T])
	PrintFormulaOfBrown(ewma, brown, alpha, T)
	var b, result float64
	b = GetTrendOfBrown(ewma, brown, alpha, T)
	for i := 0; i <= after; i++ {
		result = (2 * ewma[T]) - brown[T] + (float64(i) * b)
		fmt.Printf("%d년 예측값 : %f\n", year[T]+i, result)
		predict = append(predict, result)
	}

	return predict
}

func BrownPredict2(ewma, brown []float64, alpha float64, T, k int) (predict float64) {
	b := GetTrendOfBrown(ewma, brown, alpha, T)
	c := GetConstantOfBrown(ewma, brown, alpha, T)
	predict = c + (float64(T)+float64(k))*b
	return predict
}
