package main

import "testing"

func Test_split(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "abc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(); got != tt.want {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
