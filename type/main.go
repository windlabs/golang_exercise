package main

import "fmt"

type myInt int

func (m *myInt) string() string {
	return fmt.Sprintf("print strint %v", *m)
}

func main() {
	//var a int
	var b myInt

	// a = 1
	// b = 2
	fmt.Println(b.string())

	// var b byte
	// b = 'r'
	// fmt.Printf("%T\n", b)
}
