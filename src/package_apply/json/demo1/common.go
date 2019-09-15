package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"asd", 11}
	j, _ := json.Marshal(&person)
	fmt.Println(string(j))
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["a"] = 1
	m["b"] = "d"
	m["c"] = [1]int{1}
	j1, _ := json.Marshal(&m)
	fmt.Println(string(j1))
	var mslice []map[string]interface{}
	mslice = append(mslice, m)
	m["a"] = 2
	mslice = append(mslice, m) // 引用
	j2, _ := json.Marshal(&mslice)
	fmt.Println(string(j2))

}
