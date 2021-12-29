package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Anamial struct {
	Name string
}

type st struct {
	f1 int64
	f2 byte
}

func main() {

	var a float32 = 3.14
	reflectType(a)
	var b int = 100
	reflectType(b)
	var c Anamial
	reflectType(c)
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("type:%+v, kind:%v, value:%v.\n", t.Name(), t.Kind(), v)
}

func alignSize() {
	var s st
	fmt.Println("st")
	fmt.Println("对齐宽度", unsafe.Alignof(s))
	fmt.Println("本身大小", unsafe.Sizeof(s))
	fmt.Println("f1")
	fmt.Println("对齐宽度", unsafe.Alignof(s.f1))
	fmt.Println("本身大小", unsafe.Sizeof(s.f1))
	fmt.Println("f2")
	fmt.Println("对齐宽度", unsafe.Alignof(s.f2))
	fmt.Println("本身大小", unsafe.Sizeof(s.f2))
}
func f(x int) (_, __ int) {
	_, __ = x, x
	return
}
func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
func hello(i *int) int {
	defer func() {
		*i = 19
	}()
	return *i
}

func closeChan() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
}
func deferF() {
	f := func() { fmt.Print("A") }
	defer f()
	f = func() { fmt.Print("B") }
	defer f()
}
func uint() {
	x := []byte{}
	fmt.Printf("%+v %T\n", x, x)
}
func abc() {
	//funcs := map[string]interface{}{"foo":foo}
	//
	//for _, f := range funcs{
	////   ff := reflect.ValueOf(f)
	////   //params := make([]reflect.Value, 1)
	////   //params[0] = reflect.ValueOf(1)
	////   ff.Call([]reflect.Value{reflect.ValueOf(1)})
	////
	////}
	//
	//m := make(map[int]int,1000)
	//go func() {
	//	for i:=0;i<100;i++{
	//		m[i] = 1
	//	}
	//
	//}()
	//
	//go func() {
	//	for i:=0;i<100;i++{
	//		m[i] = 1
	//	}
	//}()
	//
	//time.Sleep(1*time.Second)
	//fmt.Println(m)
}
