package main

import (
	"fmt"
	"reflect"
)

func testReflect(val interface{}) {
	type_ := reflect.TypeOf(val)
	fmt.Printf("%v ", type_)
	fmt.Printf("%v ", type_.Name())
	fmt.Printf("%v ", type_.Kind())
}

func testReflect_Person(val interface{}) {
	type_ := reflect.TypeOf(val)

	fields_num := type_.NumField()
	for i := 0; i < fields_num; i++ {
		field := type_.Field(i)
		fmt.Printf("%v tag:%v \n", field, field.Tag.Get("a"))
	}
}

type person struct {
	Name string `a:"1"`
	age  int
}

func (p person) Setname() {

}
func main() {

	testReflect_Person(person{})
}
