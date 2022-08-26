package main

import (
	"reflect"
	"testing"
	"windlabs.com/base/bit_permission/enum"
)

func Test_permissionFmt(t *testing.T) {
	tests := []struct {
		name string
		want []enum.PermissionType
	}{
		{
			"first_case",
			[]enum.PermissionType{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permissionFmt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permissionFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
