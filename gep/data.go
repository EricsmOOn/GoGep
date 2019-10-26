package gep

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Data struct {
	TermVarSet []float64
	Result     float64
}

var sds []Data

var tds []Data

func ReadSampleData() []Data {
	return sds
}

func ReadTestData() []Data {
	return tds
}

func InitSampleData() {
	sds = make([]Data, 0)
	csvFile := sampleDataScanner()
	varSetNum := len(csvFile[0]) - 1
	var sd Data
	for _, s := range csvFile {
		sd = Data{make([]float64, varSetNum), 0.0}
		for i := 0; i < varSetNum; i++ {
			sd.TermVarSet[i], _ = strconv.ParseFloat(s[i], 64)
		}
		sd.Result, _ = strconv.ParseFloat(s[varSetNum], 64)
		sds = append(sds, sd)
	}

	if TenCheck {
		lens := len(sds)
		r := R.Intn(lens / 10)
		tds = make([]Data, 0)
		tds = append(tds, sds[r:r+10]...)
		sds = append(sds[:r], sds[r+11:]...)
	}

	TermSet = TermSetAll[:GetVarSetNum()]
	ResultRang = SelectRang*float64(GetSampleDataNum()) - Precision
	ResultSampleAvg = GetSampleResultAvg()
}

func sampleDataScanner() [][]string {
	csvFile, err := os.Open("./" + CsvSampleFileName)
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
	data := ReadSampleData()
	return len(data[0].TermVarSet)
}

func GetSampleDataNum() int {
	data := ReadSampleData()
	return len(data)
}

func GetTestDataNum() int {
	data := ReadTestData()
	return len(data)
}

func GetSampleResultAvg() float64 {
	data := ReadSampleData()
	sum := 0.0
	for _, d := range data {
		sum += d.Result
	}
	return sum / float64(GetSampleDataNum())
}
