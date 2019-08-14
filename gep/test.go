package gep

type TestData struct {
	TermVarSet []byte
	Result int
}

func ReadTestData() []TestData {
	return []TestData{
		TestData{[]byte{'1','2'},-9},
		TestData{[]byte{'4','3'},35},
		TestData{[]byte{'5','2'},56},
		TestData{[]byte{'7','6'},-65},
		TestData{[]byte{'3','2'},-5},
		TestData{[]byte{'3','1'},20},
		TestData{[]byte{'4','2'},0}}
}
