package main

import (
	"fmt"
	"strconv"
)

func main() {
	strConv()
}

func strConv() {
	i := 97
	s := string(i)
	s2 := strconv.Itoa(i)

	fmt.Printf("s value:%v, s2 value:%v\n", s, s2)

	q := strconv.QuoteToASCII("Hello, 世界")
	fmt.Printf("%v", q)

	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	k, err := strconv.ParseInt("-42", 10, 64)

	u, err := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, k, u, err)
}
