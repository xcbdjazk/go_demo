package main

import (
	"fmt"
	"reflect"
)

type P struct {
	Name string
	Age  int
}
type myint *int

func (this *P) GetName() string {
	return this.Name
}

func main() {
	i := 1
	a := myint(&i)
	fmt.Println(*a)
	p := P{}
	fmt.Println(reflect.TypeOf(p))
	// reflect.TypeOf 获取变量的类型， 返回 reflect.Type
	fmt.Println(reflect.ValueOf(p).Interface())
	// reflect.ValueOf 获取变量的类型， 返回 reflect.Value ---> 是一个结构体
}
