package main

import "fmt"

type I interface {
	run() int
}

type A struct {
}

type C struct {
	I
}

func (a A) run() int {
	fmt.Println("a run")
	return 123
}

type B struct {
	A
}

func (a B) run() int {
	fmt.Println("b run")
	return 123
}

func main() {
	var a I
	var b I = B{}
	a = A{}
	_ = a.run()
	b.(B).A.run()
	var c B
	c.A.run()

	var a1 = [...]int{1, 2, 3}
	_ = func(ddd [3]int) error {
		ddd[1] = 0
		return nil
	}(a1)
	fmt.Println(a1)
}
