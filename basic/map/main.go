package main

import "fmt"

func main() {
	create()
}

func create() {
	var m = make(map[string]string)
	m["a"] = "a"
	m["a"] = "a"
	fmt.Println(m)
	var m1 = map[string]string{
		"a": "b",
		"c": "b",
		"d": "b",
	}
	m1["a"] = "a"
	a, _ := m1["x"]
	fmt.Printf("类型 %T 值 %v \n", a, a)
	delete(m1, "b")
}
