package gep

type TestData struct {
	TermVarSet []float64
	Result     float64
}

func ReadTestData() []TestData {
	return []TestData{
		TestData{[]float64{6.9408}, 44.91},
		TestData{[]float64{-7.8664}, 7.341},
		TestData{[]float64{-2.7861}, -4.477},
		TestData{[]float64{-5.0944}, -2.307},
		TestData{[]float64{9.4895}, 73.494},
		TestData{[]float64{-9.6197}, 17.41},
		TestData{[]float64{-9.4145}, 16.073},
		TestData{[]float64{-0.1432}, -0.419},
		TestData{[]float64{0.9107}, 3.147},
		TestData{[]float64{2.1762}, 8.897}}
}

func GetDataNum() int {
	data := ReadTestData()
	return len(data)
}
