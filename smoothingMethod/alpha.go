package smoothingMethod

import (
	"fmt"
	"math"

	"github.com/evaluation"
)

func GetAlpha(observed []float64, year map[int]int, T int) (alpha float64) {

	var min float64 = math.MaxFloat64
	findAlpha := make(map[float64]float64)
	predicted := make([]float64, len(observed))
	for alpha = 0.01; alpha < 1; alpha += 0.01 {
		ewma := EWMAs(observed, alpha)
		brown := EWMAs(ewma, alpha)
		j := 0
		for i := -(year[T]); i < 3; i++ {
			predicted[j] = BrownPredict2(ewma, brown, alpha, year[T], i+2)
			j++
		}
		temp := evaluation.MSE(observed, predicted)
		fmt.Printf("alpha가 %f일 때, MSE : %f\n", alpha, temp)
		if min > temp {
			min = temp
		}
		findAlpha[temp] = alpha
	}

	return findAlpha[min]
}
