package main

import "fmt"

func main() {

	var b [4]int
	doSomeThing(b[:])
	fmt.Println(b)
}
func doSomeThing(b []int) {
	b[0] = 4
}
