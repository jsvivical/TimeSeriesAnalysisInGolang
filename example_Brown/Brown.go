package main

import (
	"fmt"
	"log"
	"os"

	"github.com/evaluation"
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

	//최적 alpha 구하기
	//먼저 훈련용, 테스트용 데이터 생성함
	var training, test []float64
	training, test = evaluation.MakeTrainingData(DF.Col("case").Float())
	alpha := smoothingMethod.GetAlphaOfBrown(training, test, len(training)-1)
	//최적 알파 확인
	fmt.Println(alpha)
	//연도 저장
	year, err := DF.Col("\ufeffyear").Int()
	if err != nil {
		log.Fatal(err)
	}
	//ewma를 저장할 변수 선언
	var ewma []float64
	ewma = smoothingMethod.EWMAs(DF.Col("case").Float(), 0.2) //alpha는 임의의 값으로 하였음

	for i, v := range ewma {
		fmt.Println(year[i], " : ", v)
	}
	fmt.Println("\n\n\n")

	brown := smoothingMethod.Brown(ewma, 0.2)
	for i, v := range brown {
		fmt.Println(year[i], ":", v)
	}
	smoothingMethod.PrintFormulaOfBrown(ewma, brown, 0.2, DF.Nrow()-4)

}
