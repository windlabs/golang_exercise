package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"time"
	// "os"
	// "os/signal"
)

const (
	num      = 100000
	rangeNum = 100000
)

func main() {
	randSeed := rand.New(rand.NewSource(time.Now().Unix()+ time.Now().UnixNano()))
	arr := make([]int, num)

	for i:=0;i<num;i++{
		arr[i] = randSeed.Intn(rangeNum)
	}

	t := time.Now()
    originalArr := make([]int , num)
	copy(originalArr, arr)

	//冒泡排序
	//bubble(arr)
	// 选择排序
	//selectSort(arr)
	// 插入排序
	//insertSort(arr)
	//希尔排序
	shellSort(arr)
	//快速排序
	//quickSort(arr)
	// 归并排序
	//mergeSort(arr)
	// 堆排序
	heapSort(arr)

	//fmt.Println(arr)
	fmt.Println(time.Since(t))
	sort.Ints(originalArr)
	fmt.Println(arr[:10], originalArr[:10], isSame(arr, originalArr))

	//等待退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	fmt.Println("Receive ctrl-c")
}

func isSame(s1 []int, s2 []int) bool {
	if len(s1) != len(s2){
		return false
	}
	if (s1 != nil) != (s2 != nil) {
		return false
	}

	for i,v := range s1{
		if v != s2[i]{
			return false
		}
	}

	return true
}

//冒泡排序
func bubble(arr []int){
	n := len(arr)

	f := false
	for i:=n-1;i>0;i--{
		for j:=0;j<i;j++{
			if arr[j]> arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
				f = true
			}
		}
		if f == false{
			break
		}
	}
}

//选择排序
func selectSort(arr []int){
	n := len(arr)

	for i:=0;i<n;i++{
		min := i
		for j:=i;j<n;j++{
			if arr[j] < arr[min]{
				min = j
			}
		}
		if min != i {
			arr[min] , arr[i] = arr[i], arr[min]
		}
	}
}


//插入排序
func insertSort(arr []int ){
	n := len(arr)

	for i:=1; i<n;i++{
		for j:=i;j>0;j--{
			if arr[j-1] > arr[j]{
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

//希尔排序
func shellSort(arr []int){
	n := len(arr)

	for step := n/2; step>0; step /=2 {
		for i:= step;i< n;i++{
			for j:=i; j>=step ; j -= step {
				if arr[j]< arr[j-step] {
					arr[j] , arr[j-step] = arr[j-step], arr[j]
				}

			}
		}
	}
}



//快速排序
func quickSort(arr []int){
	quick(arr, 0, len(arr)-1)
}
func quick(arr []int, l int, r int){
	if l < r {
		pivot := arr[r]
		i := l-1
		for j:=l;j<r;j++{
			if arr[j] <= pivot{
				i++
				arr[i],arr[j] = arr[j], arr[i]
			}
		}

		i++
		arr[r], arr[i] = arr[i], arr[r]
		quick(arr, l, i-1)
		quick(arr, i+1, r)
	}

}

//归并排序
func mergeSort(arr []int){
	n := len(arr)
    mergeSortExec(arr, 0, n-1)
}
func mergeSortExec(arr []int, l int , r int){
	if l < r{
		mid := l+(r-l)/2
		mergeSortExec(arr, l, mid)
		mergeSortExec(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, l,mid,r int){
	list := make([]int, r-l+1)
	i,j,k := l, mid+1, 0

	for ;i<= mid && j<=r;k++{
		if arr[i] < arr[j] {
			list[k] = arr[i]
			i++
		}else{
			list[k] = arr[j]
			j++
		}
	}

	for ;i <= mid;i++ {
		list[k] = arr[i]
		k++
	}
	for ;j<=r;j++{
		list[k] = arr[j]
		k++
	}

	for m:=0;m<k;m++{
		arr[l] = list[m]
		l++
	}
}

//堆排序
func heapSort(arr []int){
	n := len(arr)
	makeMinHeap(arr)

	for i:=n-1; i>0;i--{
		arr[0], arr[i] = arr[i], arr[0]
		heapFloat(arr, 0, i)
	}
}
//func heapSort(arr []int){
//	n := len(arr)
//    makeMaxHeap(arr)
//
//	for i:=n-1; i>0;i--{
//		arr[0], arr[i] = arr[i], arr[0]
//		heapSink(arr, 0, i)
//	}
//}

func makeMaxHeap(arr []int){
	n := len(arr)
	for i:= n/2-1;i>=0;i--{
		heapSink(arr, i, n)
	}
}

func makeMinHeap(arr []int){
	n := len(arr)
	for i:= n/2-1;i>=0;i--{
		heapFloat(arr, i, n)
	}
}
func heapSink(arr []int, i int, n int){
	j := 2*i+1
	for j < n {
		if j+1 < n && arr[j+1] > arr[j] {
			j = j+1
		}

		if arr[i] > arr[j]{
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i = j
		j = 2*j+1
	}

}

func heapFloat(arr []int, i int, n int){
	j := 2*i+1
	for j < n {
		if j+1 < n && arr[j+1] < arr[j] {
			j = j+1
		}

		if arr[i] < arr[j]{
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i, j = j , 2*j+1
	}

}








func heapSort1(arr []int){
	n := len(arr)

	for i:= (n-1)/2; i>=0;i--{
       heapSink1(arr, i, n)
	}
	for i := n-1; i>0;i--{
		arr[0], arr[i] = arr[i], arr[0]
		heapSink1(arr, 0, i)
	}
}

func heapSink1(arr []int, i,n int){
	j := 2*i+1
	for j<n {
		if j+1 <n && arr[j+1] > arr[j]{
			j++
		}

		if arr[i] > arr[j]{
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i = j
		j = 2*j+1
	}
}