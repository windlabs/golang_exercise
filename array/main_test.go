package main

import (
	"reflect"
	"testing"
)

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion(tt.args.list, tt.args.l, tt.args.r, tt.args.k, tt.args.ans); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
