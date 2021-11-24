package main

import "fmt"

func main()  {
   s_append()
}

func s_append(){
	numSlice := make([]int, 2000)
	numSlice2 := make([]int, 550)
	fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	numSlice = append(numSlice, numSlice2...)
	fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	for i :=0; i<10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	}
}