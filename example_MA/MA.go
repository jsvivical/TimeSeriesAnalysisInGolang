package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/smoothingMethod"
)

func main() {
	//이동평균구하기
	//먼저 사용할 데이터 파일을 연다
	f, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//데이터 프레임을 생성
	dataDF := dataframe.ReadCSV(f)
	fmt.Println(dataDF)
	//연도를 인덱스로 바꿀 map 생성(T값으로 사용됨)
	year := make(map[int]int)
	data, err := dataDF.Col("\ufeff연도").Int()
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range data {
		year[v] = i
	}
	fmt.Println(year)
	//이동평균 값을 저장할 float64슬라이스 생성
	var ma []float64
	//MovingAverages 함수호출
	ma = smoothingMethod.MovingAverages(dataDF.Col("건수").Float(), 3)

	for i, v := range ma {
		fmt.Printf("[%d] %f\n", data[i], v)
	} //출력

}
