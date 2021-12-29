package main

import (
	"flag"
	"fmt"
)

func main() {
	uid := flag.Int("uid", 0, "请输入用户标识")
	username := flag.String("uname", "", "请输入用户名字")
	flag.Parse()
	fmt.Printf("uid:%v, name:%v.\n", *uid, *username)
	fmt.Printf("arg0,%v, args:%v, Num args:%v, nFlag:%v.", flag.Arg(0), flag.Args(), flag.NArg(), flag.NFlag())
}
