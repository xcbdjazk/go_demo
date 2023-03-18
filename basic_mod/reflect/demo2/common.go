package main

import (
	"fmt"
	"reflect"
)

func testReflect(val interface{}) {
	value := reflect.ValueOf(val)
	fmt.Println(value.Int())
	// ValueOf 转 interface{}
	valI := value.Interface()
	fmt.Println(valI)
	val2 := valI.(int)
	fmt.Println(val2)

}
func main() {
	testReflect(1)
}
