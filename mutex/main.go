package main

import (
	"fmt"
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
		fmt.Println(x)
		lock.Unlock()
	}
}

func main() {
	wg.Add(1)
	lock.Lock()
	go add()
	//fmt.Println("main")
	lock.Unlock()

	// go add()
	// go add()
	wg.Wait()
	//fmt.Println(x)
}
