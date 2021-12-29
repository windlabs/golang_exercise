package main

import (
	"fmt"
	"time"
)

func main() {
	//    timetick()
	// timeFormat()
	// fmt.Println(time.Now().UnixMicro())
	// time.Sleep(1 * time.Second)
	// fmt.Println("abc")
	// fmt.Println(time.Now().UnixMicro())
	timeZone()
}

func timeZone() {
	now := time.Now()

	fmt.Println(now)
	timeParse, err := time.Parse("2006-01-02 15:04:05", "2021-12-04 14:04:00")
	if err != nil {
		fmt.Printf("Time Parse failed, err:%v.", err)
		return
	}
	fmt.Println(timeParse)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("Load location failed, err:%v.", err)
		return
	}

	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-04 14:10:00", loc)
	if err != nil {
		fmt.Printf("Parse time failed, err:%v.", err)
		return
	}

	fmt.Println(timeObj)

	td := timeObj.Sub(time.Now())
	fmt.Println(td)

}
func timeFormat() {
	fmt.Println(time.Now().Unix(), time.Now().UnixNano())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006/01/03 15:04:05.000 PM"))
	timeObj, err := time.Parse("2006-01-02 15:04", "2021-12-03 13:00")
	if err != nil {
		fmt.Printf("Time parse failed, err:%v.", err)
		return
	}

	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

}
func timetick() {
	tick := time.Tick(time.Second * 1)
	after := time.After(time.Second * 2)

	for {
		select {
		case x := <-tick:
			fmt.Println("tick", x)
		case y := <-after:
			fmt.Println("after:", y)
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
