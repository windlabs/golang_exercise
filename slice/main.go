package main

import "fmt"

func main()  {
   s_append()
}

func s_append(){
	//numSlice := make([]int, 2000)
	//numSlice2 := make([]int, 550)
	//fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	//numSlice = append(numSlice, numSlice2...)
	//fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	//for i :=0; i<10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("len:%d cap:%d pth:%p\n", len(numSlice), cap(numSlice), numSlice)
	//}
	x1 := [...]int{1,3,5}
	s1 := x1[:1]
	s2 := x1[:2]
	ss(s2)
	fmt.Printf("%p , %p, %p,  %v,%v \n", s1, s2, x1,s1,s2)
}

func ss(s []int){
	//s = []int{1,2,5}
	s[0] = 5
	fmt.Println(s)
}