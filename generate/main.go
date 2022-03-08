package main

import (
	"fmt"
	"windlabs.com/base/generate/errcode"
)

func main(){
	code := errcode.ERR_CODE_TIMEOUT
	fmt.Println(int(code), code)
}