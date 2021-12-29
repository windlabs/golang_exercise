package main

import (
	"sync"
)

var (
	x    int = 0
	lock sync.Mutex
	wg   sync.WaitGroup
)

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
}
