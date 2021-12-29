package main

import "fmt"

const (
	DEBUG int = iota + 1
	TRACE
	INFO
	WARNING = "warning"
	ERROR
	ERROR2
	FATAL = iota
	TEST
)

func main() {
	fmt.Println(DEBUG, TRACE, INFO, WARNING, ERROR, ERROR2, FATAL, TEST)
}
