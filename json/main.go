package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"liu","age":16}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Printf("%#v", p)
	fmt.Println(p.Name)
}
