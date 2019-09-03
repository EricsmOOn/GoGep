package timer

import (
	"fmt"
	"runtime"
	"time"
)

var funcName = make(map[string]float64)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	// 获取函数名称
	fn := runtime.FuncForPC(pc).Name()
	return fn
}

func TimeCount() func() {
	name := getFunctionName()
	start := time.Now()
	return func() {
		tc := time.Since(start)
		value := funcName[name]
		if value != 0 {
			funcName[name] = value + tc.Seconds()
		} else {
			funcName[name] = tc.Seconds()
		}
	}
}

func PrintTimer() {
	for k, v := range funcName {
		fmt.Printf("%10s - %.5fs\n", k, v)
	}
}
