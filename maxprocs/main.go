package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A:%v\n", i)
	}
}

func b() {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		fmt.Printf("B:%v\n", j)
	}
}
func main() {
	wg.Add(2)
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	wg.Wait()
}
