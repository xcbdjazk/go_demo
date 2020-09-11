package main

import "fmt"

func f1(i int) {
	i = 200
}
func f2(i *int) {
	*i = 1200
}

func zhizheng() {

	i := 20
	f1(i)
	fmt.Println(i)
	f2(&i)
	fmt.Println(i)
}

func newUse() {
	// new 给指针变量分配存储空间
	var a = new(int)
	*a = 200
	fmt.Println(*a)
}

func main() {
	newUse()
}
