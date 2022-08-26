package main

import (
	"fmt"
	"time"
)

func main() {
	func1()
	//func2()
}
func func1() {
	go func() {
		for {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				proc()
			}()
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}

func func2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from f.", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%#v", i))
	}

	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g.", i)
	g(i + 1)
}
