package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

func main() {
	wg2()

	//limitChan(10)
	//for i :=0;i<5;i++{
	//	rand1 := rander.Int()
	//	rand2 := rander.Intn(10)
	//	fmt.Println(rand1, rand2)
	//}
	//// os single
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Kill)
	//<-c
	//fmt.Println("receive ctrl+c")

	//wg.Add(3)
	//ch1 := make(chan int, 50)
	//ch2 := make(chan int, 50)
	//go f1(ch1)
	//go f2(ch1,ch2)
	//go f2(ch1,ch2)
	//
	//for y := range ch2 {
	//	fmt.Printf("y: %+v\n", y )
	//}
	//wg.Wait()

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

// func f2(ch1 chan int, ch2 chan int) {
// 	defer wg.Done()
// 	for x := range ch1 {
// 		ch2 <- x * x
// 	}
// 	once.Do(func() { close(ch2) })
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

//func wg(){
//	list1 := []string{"a","b","c","d"}
//	list2 := []int{1,2,3,4}
//	ch1 := make(chan bool)
//	ch2 := make(chan bool)
//	exit := make(chan bool)
//
//	go func() {
//		for i:=0; i<len(list1);i++{
//			fmt.Print(list1[i])
//			ch2<-true
//			<-ch1
//		}
//		exit <-true
//	}()
//
//	go func() {
//		for j:=0;j<len(list2);j++{
//			<-ch2
//			fmt.Print(list2[j])
//			ch1<-true
//		}
//	}()
//    <-exit
//	close(ch1)
//	close(ch2)
//}
func wg2() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan bool)
	ch2 := make(chan bool)

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
	close(ch1)
	close(ch2)
}
