package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x  int32
	wg sync.WaitGroup
)

func add() {
	// x++
	atomic.AddInt32(&x, 1)
	wg.Done()
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
