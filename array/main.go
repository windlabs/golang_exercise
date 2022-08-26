package main

import (
	"fmt"
)

func main() {
	//ans := recursion([]int{1, 2, 3}, 0, 2, 2, []int{})
	//fmt.Println(ans)
	//testSlice := []int{1,2,3}
	//fmt.Printf("原始内存地址：%p\n", testSlice)
	//test(testSlice)
	//fmt.Println(testSlice)
	//for i := 0; i < 3; i++ {
	//	now := time.Now()
	//	fmt.Println("____________", now.Format("2006-01-02 13:04:05"))
	//	time.Sleep(time.Duration(3) * time.Second)
	//	now = time.Now().AddDate(0, 0, -6)
	//	fmt.Println("____________", now.Format("2006-01-02 13:04:05"))
	//}
	ipList := []int{1, 2, 4}
	var ipScene int
	for _,v := range ipList{
		ipScene |= v
	}
	fmt.Println(ipScene,1&ipScene, 2&ipScene)
}

func test(testSlice []int) {
	fmt.Printf("函数接收内存地址：%p\n", testSlice)
	testSlice = []int{}
	fmt.Printf("新赋值后内存地址:%p\n", testSlice)
}

// 递归
func recursion(list []int, l int, r int, k int, ans []int) []int {
	if l < r {
		pivot := list[r]
		i := l
		for j := l; j < r; j++ {
			if list[j] < pivot {
				list[i], list[j] = list[j], list[i]
				i++
			}
		}
		list[i], list[r] = list[r], list[i]
		bigLen := r - i + 1
		if bigLen <= k {
			for n := i; n <= r; n++ {
				ans = append(ans, list[n])
			}
		}

		if bigLen > k {
			return recursion(list, i, r, k, ans)
		}
		if bigLen < k {
			return recursion(list, l, i-1, k-bigLen, ans)
		}
	}
	return ans
}
