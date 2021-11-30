package main

import "fmt"

func main() {
	var a = new(int)
	var b = make([]int, 1)

	fmt.Printf("a type: %T, value:%v; b type%T, value%v \n", a, *a, b, b)
}
