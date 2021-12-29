package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("一条测试日志")
		time.Sleep(1 * time.Second)
	}
}
