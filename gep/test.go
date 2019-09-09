package gep

func GetTestResult() []float64 {
	var testResults = make([]float64, 0)
	tds := ReadTestData()
	for _, td := range tds {
		testResults = append(testResults, td.Result)
	}
	return testResults
}

func GetSampleResult() []float64 {
	var sampleResults = make([]float64, 0)
	sds := ReadSampleData()
	for _, sd := range sds {
		sampleResults = append(sampleResults, sd.Result)
	}
	return sampleResults
}

func GetPredictTestResult(easyEquation [][]byte) []float64 {
	var testResults = make([]float64, 0)
	testData := ReadTestData()
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, _ := calculatePerFitOpt(easyEquation, td.TermVarSet)
		testResults = append(testResults, result)
	}
	return testResults
}

func GetPredictSampleResult(easyEquation [][]byte) []float64 {
	var sampleResults = make([]float64, 0)
	sampleData := ReadSampleData()
	for _, sd := range sampleData {
		//逐个读入测试数据
		//求表达 result
		result, _ := calculatePerFitOpt(easyEquation, sd.TermVarSet)
		sampleResults = append(sampleResults, result)
	}
	return sampleResults
}
