package main

import "testing"

func Test_jq(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jq(); got != tt.want {
				t.Errorf("jq() = %v, want %v", got, tt.want)
			}
		})
	}
}
