package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	lock sync.Mutex
	m    = sync.Map{}
)

func split() string {
	return "abc"
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered err:%v", err)
		}
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			m.Store(n, n)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 11; i++ {
		v, ok := m.Load(i)
		if !ok {
			fmt.Println(ok)
		}
		fmt.Printf("%v\n", v)
	}

}
