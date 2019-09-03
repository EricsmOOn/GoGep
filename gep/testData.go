package gep

type TestData struct {
	TermVarSet []float64
	Result     float64
}

var sunspots = []float64{
	101, 82, 66, 35, 31, 7, 20, 92,
	154, 125, 85, 68, 38, 23, 10, 24,
	83, 132, 131, 118, 90, 67, 60, 47,
	41, 21, 16, 6, 4, 7, 14, 34,
	45, 43, 48, 42, 28, 10, 8, 2,
	0, 1, 5, 12, 14, 35, 46, 41,
	30, 24, 16, 7, 4, 2, 8, 17,
	36, 50, 62, 67, 71, 48, 28, 8,
	13, 57, 122, 138, 103, 86, 63, 37,
	24, 11, 15, 40, 62, 98, 124, 96,
	66, 64, 54, 39, 21, 7, 4, 23,
	55, 94, 96, 77, 59, 44, 47, 30,
	16, 7, 37, 74}

//func ReadTestData() []TestData {
//	return []TestData{
//		//TestData{[]float64{6.9408}, 44.91},
//		//TestData{[]float64{-7.8664}, 7.341},
//		//TestData{[]float64{-2.7861}, -4.477},
//		//TestData{[]float64{-5.0944}, -2.307},
//		//TestData{[]float64{9.4895}, 73.494},
//		//TestData{[]float64{-9.6197}, 17.41},
//		//TestData{[]float64{-9.4145}, 16.073},
//		//TestData{[]float64{-0.1432}, -0.419},
//		//TestData{[]float64{0.9107}, 3.147},
//		//TestData{[]float64{2.1762}, 8.897}}
//
//		TestData{[]float64{-2.121070}, -4.113741},
//		TestData{[]float64{-2.508795}, -4.379359},
//		TestData{[]float64{3.546891}, 16.930894},
//		TestData{[]float64{-0.498116}, -1.370288},
//		TestData{[]float64{7.957594}, 55.534436},
//		TestData{[]float64{-3.483753}, -4.382991},
//		TestData{[]float64{-7.526398}, 5.744140},
//		TestData{[]float64{0.945305}, 3.282716},
//		TestData{[]float64{6.979210}, 45.292315},
//		TestData{[]float64{-9.322047}, 15.484142}}
//
//	//TestData{[]float64{5.7695}, 232.107},
//	//TestData{[]float64{-4.1206}, -56.1063},
//	//TestData{[]float64{4.652}, 127.968},
//	//TestData{[]float64{-5.6193}, -150.481},
//	//TestData{[]float64{-3.2971}, -27.2686},
//	//TestData{[]float64{-0.0599}, 0.943473},
//	//TestData{[]float64{1.1835}, 5.24187},
//	//TestData{[]float64{-8.2814}, -506.651},
//	//TestData{[]float64{4.4342}, 112.282},
//	//TestData{[]float64{4.1843}, 95.9529}}
//
//}

//太阳黑子
func ReadTestData() []TestData {
	td := make([]TestData, 0)
	for i := 0; i < 90; i++ {
		td = append(td, TestData{sunspots[i : i+10], sunspots[i+10]})
	}
	return td
}

func GetDataNum() int {
	//太阳黑子
	//return 90
	data := ReadTestData()
	return len(data)
}
