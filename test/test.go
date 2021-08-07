package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/evaluation"
	"github.com/go-gota/gota/dataframe"
	"github.com/smoothingMethod"
)

type ts struct {
	idx  int
	time string
	data float64
}

func main() {
	var data []float64
	var timeSeries []ts
	var i, N int
	var year []string
	var record []string
	var temp float64

	search := make(map[string]int)
	i = 0
	f, err := os.Open("sample2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			i++
			continue
		}
		temp, err = strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(temp)
		year = append(year, record[0]) //년도
		data = append(data, temp)      //값
		search[record[0]] = i - 1
		timeSeries = append(timeSeries, ts{(i - 1), record[0], temp})
		i++
	}
	fmt.Println(search)
	for _, v := range timeSeries {
		fmt.Println(v)
	}
	N, _ = strconv.Atoi(os.Args[1])

	MAResult := smoothingMethod.MovingAverages(data, N)
	fmt.Println(MAResult)
	for i, v := range MAResult {
		fmt.Printf("[%v] %f\n", year[i], v)
	}
	DMAResult := smoothingMethod.DoubleMovingAverages(MAResult, N)
	for i, v := range DMAResult {
		fmt.Printf("[%v] %f\n", year[i], v)
	}

	tempT := os.Args[2]

	smoothingMethod.PrintFormula(MAResult, DMAResult, N, search[tempT])

	predict := smoothingMethod.Predict(MAResult, DMAResult, N, search[tempT], 3)
	fmt.Println(predict)

	f2, err := os.Open("sample3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	DF := dataframe.ReadCSV(f2)

	var training, test []float64
	training, test = evaluation.MakeTrainingData(DF.Col("case").Float())
	fmt.Println("\n\ntraining")
	fmt.Println(training)
	fmt.Println("\n\ntest")
	fmt.Println(test)
	fmt.Println("\n\n")
	fmt.Println(DF)

	alpha := smoothingMethod.GetAlphaOfBrown(training, test, len(training)-1)
	fmt.Println(alpha)

	y, err := DF.Col("\ufeffyear").Int()
	if err != nil {
		log.Fatal(err)
	}
	var ewma []float64
	ewma = smoothingMethod.EWMAs(DF.Col("case").Float(), 0.2)

	for i, v := range ewma {
		fmt.Println(y[i], ":", v)
	}
	println("\n\n\n")
	brown := smoothingMethod.Brown(ewma, 0.2)
	for i, v := range brown {
		fmt.Println(y[i], ":", v)
	}
	smoothingMethod.PrintFormulaOfBrown(ewma, brown, 0.2, DF.Nrow()-4)
}
