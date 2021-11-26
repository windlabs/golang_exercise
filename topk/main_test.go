package main

import (
	"sort"
	"testing"
)
//测试循环方式
func Test_loop(t *testing.T) {
	type args struct {
		list []int
	}
	type params struct {
		name       string
		args       args
		wantFirst  int
		wantSecond int
	}
	tests := []params{
		{"loop1",args{[]int{1,2,3,4}},4,3},
		{"loop2",args{[]int{4,2,6,9,7,3,4}},9,7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirst, gotSecond := loop(tt.args.list)
			if gotFirst != tt.wantFirst {
				t.Errorf("loop() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
			}
			if gotSecond != tt.wantSecond {
				t.Errorf("loop() gotSecond = %v, want %v", gotSecond, tt.wantSecond)
			}
		})
	}
}
//测试递归方式
func Test_recursion(t *testing.T) {
	type args struct {
		list []int
		l    int
		r    int
		k    int
		ans  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"recursion1", args{[]int{2,3,1,6,4,7,8},0, 6, 2,[]int{}},[]int{7,8}},
		{"recursion2", args{[]int{2,3,10,6,4,7,8},0, 6, 2,[]int{}},[]int{10,8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion(tt.args.list, tt.args.l, tt.args.r, tt.args.k, tt.args.ans); !isSame(got, tt.want) {
				t.Errorf("recursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

//比较两个slice是否相同
func isSame(l1 []int, l2 []int) bool{
	sort.Ints(l1)
	sort.Ints(l2)
	if len(l1) != len(l2) {
		return false
	}

	if (l1 != nil ) != (l2 != nil) {
		return false
	}

	for i,v := range l1 {
		if v != l2[i]{
			return false
		}
	}
	return true
}