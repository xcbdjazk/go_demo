package main

import "fmt"

func main() {

	var a interface{}
	a = []int{1, 2, 3}
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("no")
	}

}
