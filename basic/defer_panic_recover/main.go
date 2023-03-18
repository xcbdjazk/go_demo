package main

import "fmt"

func f1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var i = 0
	var b = 1
	c := b / i
	fmt.Println(c)
}
func f2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic([]int{1, 2, 3})
}

func main() {
	f2()
}
