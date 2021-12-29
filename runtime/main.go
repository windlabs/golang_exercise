package main

import (
	"fmt"
	"runtime"
)

func runtimeFunc() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime caller failed")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("func name %#v \n", funcName)

	fmt.Printf("file %v\n", file)
	fmt.Printf("line %v \n", line)
}

func main() {
	runtimeFunc()
}
