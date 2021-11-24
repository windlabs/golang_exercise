package main

import (
	"fmt"
	"time"
)

func main(){
   timetick()
}

func timetick(){
	tick := time.Tick( time.Second * 1)
	after := time.After(time.Second * 2)

	for {
		select {
		case x := <-tick:
			fmt.Println("tick", x)
		case y := <- after:
			fmt.Println("after:", y)
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
