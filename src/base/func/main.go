package main

import "fmt"

func f1(d int, a ...interface{}) {
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Printf("%#v, %T ", a, a)
}

func main() {
	var a = []interface{}{
		"dd", 1, 'd',
	}
	f1(1, a[:3]...)
}
