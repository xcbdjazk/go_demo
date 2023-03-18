package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var b [4]int
	c := b[:]

	doSomeThing(c)

	ptr := unsafe.Pointer(&c)
	fmt.Println(ptr)
	fmt.Println(b)
	c = append(c, 1, 2, 3, 4)
	fmt.Println(&c)
	fmt.Println(unsafe.Pointer(&c))
}
func doSomeThing(b []int) {
	b[0] = 4
}
