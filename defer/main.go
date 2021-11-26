package main

import "fmt"

func main(){
	foo := 100
	defer func() {fmt.Println(foo)}()
	defer fmt.Println(foo)

	foo = 200
	defer func(foo int) {fmt.Println(foo)}(foo)

	foo = 300
}
