package main

import "testing"

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			split()
		})
	}
}
