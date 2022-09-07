package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	//pb := &x.b
	//pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	pb := (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	*pb = 255
	fmt.Println(x.b)
}
