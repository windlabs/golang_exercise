package main

import (
	"fmt"
)

type bar struct {
	Name string
}

func main() {
	b := new(bar)
	bb := new(*bar)
	fmt.Println(bb)
	bb = &b

	fmt.Println(bb)
	fmt.Println(b)
}
