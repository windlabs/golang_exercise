package main

import "fmt"

func main() {
	ans := recursion([]int{1, 2, 3}, 0, 2, 2, []int{})
	fmt.Println(ans)
}

//递归
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
