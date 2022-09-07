package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"w.src.corp.qihoo.net/pubsec/go-library/sync/goroutine"
)

var (
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
	wg     = sync.WaitGroup{}
	once   sync.Once
)

type result struct {
	Num   int64
	Sum   int64
	Index int
}

func createRand(jobChan chan<- int64) {
	for {
		jobChan <- rander.Int63()
		time.Sleep(1 * time.Second)
	}
}

func worker(jobChan <-chan int64, resultChan chan<- int64) {
	for {
		var ans int64
		x := <-jobChan
		for x > 0 {
			ans += x % 10
			x /= 10
		}
		resultChan <- ans
	}
}

func main() {
	goroutinePools, err := goroutine.InitGoroutinePools(2)
	if err != nil {
		return
	}
	_ = goroutinePools.Submit(func() {
		go func() {
			time.Sleep(10*time.Second)
			fmt.Println("--------------------")
		}()

	})
	_ = goroutinePools.Submit(func() {
		time.Sleep(5*time.Second)
		fmt.Println("====================")
	})
	goroutinePools.WaitFinish()

	fmt.Println("1111")
	// wg.Add(24)
	//jobChan := make(chan int64, 100)
	//resultChan := make(chan int64, 100)
	//go createRand(jobChan)
	//for i := 0; i < 2; i++ {
	//	go worker(jobChan, resultChan)
	//}
	//for {
	//	fmt.Println(<-resultChan)
	//}
	// wg.Wait()
	// wg2()

	//limitChan(10)
	//for i :=0;i<5;i++{
	//	rand1 := rander.Int()
	//	rand2 := rander.Intn(10)
	//	fmt.Println(rand1, rand2)
	//}

	//// os single
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Kill, os.Interrupt)
	//<-c
	//fmt.Println("receive ctrl+c")

	// wg.Add(2)
	// ch1 := make(chan int)
	// // ch2 := make(chan int)
	// go f1(ch1)
	// //go f2(ch1,ch2)
	// go f2(ch1)
	// wg.Wait()

	// jobChan := make(chan int64, 100)
	// resultChan :=make(chan result, 100)

	// go createJob(jobChan)
	// for i:=0;i<24;i++{
	// 	go worker(jobChan, resultChan,i )
	// }

	// for res := range resultChan {
	// 	fmt.Printf("goruntine:%+v, num:%+v, sum:%+v\n",res.Index, res.Num, res.Sum)
	// }

}

// func createJob(jobChan chan<- int64) {
// 	for {
// 		jobChan <- rander.Int63()
// 		//time.Sleep(time.Second * 1)
// 	}
// }

// func worker(jobChan <-chan int64, resultChan chan<- result, goIndex int) {
// 	for num := range jobChan {
// 		resultChan <- result{num, numSum(num), goIndex}
// 	}
// }

// func numSum(num int64) int64 {
// 	var ans int64
// 	for ; num > 0; num /= 10 {
// 		ans += num % 10
// 	}
// 	return ans
// }
// func f1(ch1 chan int) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		ch1 <- i
// 	}
// 	close(ch1)
// }

// func f2(ch1 chan int) {
// 	defer wg.Done()
// 	for {
// 		x := <-ch1
// 		fmt.Println(x)
// 	}
// }

// func limitChan(l int) {
// 	limitCh := make(chan bool, l)
// 	for i := 0; i < 100; i++ {
// 		go func(i int, l int) {
// 			fmt.Println(i, l)
// 			defer func() {
// 				<-limitCh
// 			}()
// 		}(i, len(limitCh))
// 		limitCh <- true
// 	}
// }

// func limitChan2(l int) {

// }

func wg1() {
	list1 := []string{"a", "b", "c", "d"}
	list2 := []int{1, 2, 3, 4}
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	exit := make(chan bool)

	go func() {
		for i := 0; i < len(list1); i++ {
			fmt.Print(list1[i])
			ch2 <- true
			<-ch1
		}
		exit <- true
	}()

	go func() {
		for j := 0; j < len(list2); j++ {
			<-ch2
			fmt.Print(list2[j])
			ch1 <- true
		}
	}()
	<-exit
	close(ch1)
	close(ch2)
}
func wg2() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	defer close(ch1)
	defer close(ch2)

	go func() {
		for i := 65; i < 91; i++ {

			fmt.Print(string(i))
			ch2 <- true
			<-ch1

		}
		wg.Done()
	}()

	go func() {
		for j := 1; j < 27; j++ {
			<-ch2
			fmt.Print(j)
			ch1 <- true
		}
		wg.Done()
	}()

	wg.Wait()
}

func wg3() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	for {
		go func() {
			for {
				<-ch1
				fmt.Println("cat")
				ch2 <- true
			}
		}()
		go func() {
			for {
				<-ch2
				fmt.Println("dog")
				ch3 <- true
			}
		}()
		go func() {
			for {
				<-ch3
				fmt.Println("fish")
				ch1 <- true
			}
		}()
		ch1 <- true
	}
}
