package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name:"1",args:args{1,2},want:2},
		{name:"2",args:args{3,2},want:3},
		{name:"2",args:args{3,1},want:3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{name:"1",args:args{"a:b:c",":"},want:[]string{"a","b","c"}},
		{"2",args{"abcdef","cd"},[]string{"ab","ef"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.str, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}


func BenchmarkSplit(b *testing.B){
	for i:=0;i<b.N; i++{
		split("a.b.c", ".")
	}
}

func BenchmarkMax(b *testing.B){
	for i:=0;i<b.N; i++{
		max(1, 2)
	}
}

func BenchmarkMaxParallel(b *testing.B){
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			max(2, 3)
		}
	})
}