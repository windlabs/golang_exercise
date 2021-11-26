package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	num = 1000
	numRange = 10000
)

func main()  {
	randSeed := rand.New(rand.NewSource(time.Now().Unix()+ time.Now().UnixNano()))
	list := make([]int, num)

	for i:=0;i<num;i++{
		list[i] = randSeed.Intn(numRange)
	}

	//loop
	first, second := loop(list)
	fmt.Printf("first:%v, second:%v\n", first, second)

	//recursion
	k := 2
	ans := []int{}
	ret := recursion(list, 0, len(list)-1, k, ans)
	fmt.Printf("ans:%v \n", ret)
}

//循环
func loop(list []int) (first, second int){
	for _, v := range list{
		if v > second {
			if v > first{
				first,second = v, first
			}else{
				second = v
			}
		}
	}
	return
}

//递归
func recursion(list []int, l int, r int, k int, ans []int) []int {
	if l < r {
		pivot := list[r]
		i := l
		for j:=l;j<r;j++{
			if list[j] < pivot {
				list[i], list[j] = list[j],list[i]
				i++
			}
		}
		list[i],list[r] = list[r], list[i]
		bigLen := r-i+1
		if bigLen <= k {
			for n :=i; n<=r;n++{
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