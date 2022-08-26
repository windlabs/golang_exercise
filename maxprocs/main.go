package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%v\n", i)
		time.Sleep(1 * time.Nanosecond)
	}
}

func b() {
	defer wg.Done()
	for j := 0; j < 10; j++ {
		fmt.Printf("B:%v\n", j)
		time.Sleep(1 * time.Nanosecond)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
