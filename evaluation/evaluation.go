package evaluation

import (
	"math"
)

func MakeTrainingData(totalData []float64) (training []float64, test []float64) {

	trainingNum := 4 * len(totalData) / 5
	testNum := len(totalData) / 5
	if trainingNum+testNum < len(totalData) {
		trainingNum++
	}

	training = totalData[0:trainingNum]
	test = totalData[trainingNum+1:]

	return training, test
}

func MSE(observed, predicted []float64) float64 {
	var MSE float64
	for i := 0; i < len(observed); i++ {
		MSE += math.Pow((observed[i]-predicted[i]), 2) / float64(len(observed))
	}
	return MSE
}

func MAE(observed, predicted []float64) float64 {
	var MAE float64
	for i := 0; i < len(observed); i++ {
		MAE += math.Abs((observed[i] - predicted[i])) / float64(len(observed))
	}
	return MAE
}
