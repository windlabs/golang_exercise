package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) string() string {
	return fmt.Sprintf("print %v", p.name)
}
func main() {
	var p1 person
	p1.name = "person one"

	// p2 := person{
	// 	"person two",
	// 	12,
	// }
	var p3 = new(person)

	fmt.Printf("p1: %#v, %T, p3:%#v, %T", p1, p1, p3, p3)

}
