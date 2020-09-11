package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Age  int
}

func (a A) aFunc() string {
	fmt.Println("a func")
	return a.Name + " func"
}

type B struct {
	Name string
	Age  int
	A
}

func (b *B) bFunc() string {

	return b.Name
}

func main() {
	var a = B{}

	a.Name = "123"
	a.Age = 123
	fmt.Printf("%#v", a)
	s := a.aFunc()
	fmt.Println(s)
	strByte, _ := json.Marshal(a)
	fmt.Println(string(strByte))
	d := `{"Name":"123","Age":1231, "str":"c"}`
	var b = make(map[string]interface{})
	err := json.Unmarshal([]byte(d), &b)
	fmt.Printf("%#v \n", err)
	fmt.Printf("%#v", b)
}
