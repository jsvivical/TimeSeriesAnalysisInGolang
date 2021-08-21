package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/smoothingMethod"
)

func main() {
	//먼저 지수평활이동평균을 구함(EWMA)
	//계산할 데이터 파일 열기
	f, err := os.Open("sample3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//데이터 프레임 생성
	DF := dataframe.ReadCSV(f)
	//데이터 확인용
	fmt.Println(DF)

	//연도 저장
	year, err := DF.Col("\ufeffyear").Int()
	if err != nil {
		log.Fatal(err)
	}
	//ewma를 저장할 변수 선언
	var ewma []float64
	ewma = smoothingMethod.EWMAs(DF.Col("case").Float(), 0.41) //alpha는 임의의 값으로 하였음

	for i, v := range ewma {
		fmt.Println(year[i], " : ", v)
	}
	fmt.Println("\n\n\n")

	brown := smoothingMethod.Brown(ewma, 0.41)
	for i, v := range brown {
		fmt.Println(year[i], ":", v)
	}
	yearToIndex := make(map[int]int)
	for i, v := range year {
		yearToIndex[v] = i
	}
	smoothingMethod.PrintFormulaOfBrown(ewma, brown, 0.41, yearToIndex[2013]) //2013년까지
	smoothingMethod.BrownPredict(ewma, brown, 0.41, year, yearToIndex[2013], 5)

	//방법2
	brown2 := smoothingMethod.EWMAs(ewma, 0.41)
	for i, v := range brown2 {
		fmt.Println(year[i], ":", v)
	}
	//한 시점의 예측식으로부터 예측한 관측값들
	var predicted []float64
	fmt.Println("2013 T : ", yearToIndex[2013])

	for i := -(yearToIndex[2013]); i <= 3; i++ {
		predicted = append(predicted, smoothingMethod.BrownPredict2(ewma, brown, 0.41, yearToIndex[2013], i+2))
	}

	fmt.Println("\n\n\n")
	for i, v := range predicted {
		fmt.Println(year[i], ":", v)
	}

	alpha := smoothingMethod.GetAlpha(DF.Col("case").Float(), yearToIndex, 2013)
	fmt.Println(alpha)

}
