package main

import (
	"fmt"
	"windlabs.com/base/bit_permission/enum"
)

func main() {
	fmt.Println(
		enum.Read,
		enum.Split(enum.Join(enum.Read, enum.Write)),
		enum.HasPermission(2, enum.Read),
		enum.PermissionType.Join(enum.PermissionRead, enum.PermissionWrite).Split(),
	)
}

func permissionFmt() []enum.PermissionType {
	return enum.PermissionType.Join(enum.PermissionRead, enum.PermissionWrite).Split()
}
