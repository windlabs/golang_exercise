package main

import (
	"fmt"
	"time"
)

func main(){
      ch1 := make(chan int, 10)
	  ch2 := make(chan int, 10)

	  go func() {
		  for i:=0; i<10; i++ {

			  ch1 <- i
			  fmt.Println("ch2", <-ch2)
		  }
	  }()
	go func() {
		for i:=0; i<10; i++ {
			ch2 <- i
			fmt.Println("ch1", <-ch1)
		}
	}()

	  time.Sleep(1* time.Second)
	//for  i:=0;i<20;i++{
	//	select {
	//	case v1:= <-ch1:
	//		fmt.Println("ch1", v1)
	//	case v2:= <-ch2:
	//		fmt.Println("ch2", v2)
	//	}
	//}


}