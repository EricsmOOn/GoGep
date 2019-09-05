package gep

import (
	"encoding/csv"
	"os"
	"strconv"
)

type TestData struct {
	TermVarSet []float64
	Result     float64
}

var tds []TestData

func ReadTestData() []TestData {
	return tds
}

func InitTestData() {
	tds = make([]TestData, 0)
	csvFile := testDataScanner()
	varSetNum := len(csvFile[0]) - 1
	var td TestData
	for _, s := range csvFile {
		td = TestData{make([]float64, varSetNum), 0.0}
		for i := 0; i < varSetNum; i++ {
			td.TermVarSet[i], _ = strconv.ParseFloat(s[i], 64)
		}
		td.Result, _ = strconv.ParseFloat(s[varSetNum], 64)
		tds = append(tds, td)
	}
	TermSet = TermSetAll[:GetVarSetNum()]
	ResultRang = SelectRang*float64(GetDataNum()) - Precision
}

func testDataScanner() [][]string {
	csvFile, err := os.Open("./" + CsvFileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return rows
}

func GetVarSetNum() int {
	data := ReadTestData()
	return len(data[0].TermVarSet)
}

func GetDataNum() int {
	data := ReadTestData()
	return len(data)
}
