package gep

type TestData struct {
	TermVarSet []float64
	Result     float64
}

func ReadTestData() []TestData {
	return []TestData{
		//TestData{[]float64{6.9408}, 44.91},
		//TestData{[]float64{-7.8664}, 7.341},
		//TestData{[]float64{-2.7861}, -4.477},
		//TestData{[]float64{-5.0944}, -2.307},
		//TestData{[]float64{9.4895}, 73.494},
		//TestData{[]float64{-9.6197}, 17.41},
		//TestData{[]float64{-9.4145}, 16.073},
		//TestData{[]float64{-0.1432}, -0.419},
		//TestData{[]float64{0.9107}, 3.147},
		//TestData{[]float64{2.1762}, 8.897}}

		TestData{[]float64{-2.121070}, -4.113741},
		TestData{[]float64{-2.508795}, -4.379359},
		TestData{[]float64{3.546891}, 16.930894},
		TestData{[]float64{-0.498116}, -1.370288},
		TestData{[]float64{7.957594}, 55.534436},
		TestData{[]float64{-3.483753}, -4.382991},
		TestData{[]float64{-7.526398}, 5.744140},
		TestData{[]float64{0.945305}, 3.282716},
		TestData{[]float64{6.979210}, 45.292315},
		TestData{[]float64{-9.322047}, 15.484142}}

	//TestData{[]float64{5.7695}, 232.107},
	//TestData{[]float64{-4.1206}, -56.1063},
	//TestData{[]float64{4.652}, 127.968},
	//TestData{[]float64{-5.6193}, -150.481},
	//TestData{[]float64{-3.2971}, -27.2686},
	//TestData{[]float64{-0.0599}, 0.943473},
	//TestData{[]float64{1.1835}, 5.24187},
	//TestData{[]float64{-8.2814}, -506.651},
	//TestData{[]float64{4.4342}, 112.282},
	//TestData{[]float64{4.1843}, 95.9529}}

}

func GetDataNum() int {
	data := ReadTestData()
	return len(data)
}
