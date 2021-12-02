package main

import "fmt"

func test() (res int) {
	res = 1
	defer func() {
		res++
	}()
	return 0
}

func main() {
	// foo := 100
	// defer func() {fmt.Println(foo)}()
	// defer fmt.Println(foo)

	// foo = 200
	// defer func(foo int) {fmt.Println(foo)}(foo)

	// foo = 300

	ret := test()
	fmt.Printf("test:%v", ret)
}
