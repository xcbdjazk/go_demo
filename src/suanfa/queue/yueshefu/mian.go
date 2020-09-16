package main

import (
	"fmt"
)

// 括号是否有效

func pop(l []int) ([]int, int) {
	u := l[0]
	l = l[1:]
	return l, u
}
func add(l []int, u int) []int {
	return append(l, u)
}

func test(user, m int) {
	var l []int

	for i := 1; i <= user; i++ {
		l = append(l, i)
	}
	k := 2
	l = append(l[k:], l[:k]...)
	fmt.Println(l)
	//index := 0
	start := 1
	for {
		var pu int
		if len(l) == 0 {
			break
		}
		l, pu = pop(l)
		//time.Sleep(time.Second)
		if start < m {
			start++
			l = add(l, pu)
		} else {
			fmt.Println("pop user", pu)
			start = 1
		}
	}

}

func t() {
	l := []int{1, 2, 3}

	for {
		if len(l) == 0 {
			break
		}
		l, _ := pop(l)
		fmt.Println(l)
	}
	fmt.Println(l)
}
func init() {
	test(10, 3)
	//t()
}

func testPop(l []int) ([]int, bool) {
	l = l[1:]
	return l, false
}

func main() {
	//l := []int{1,2,3}
	//l, a := testPop(l)
	//fmt.Println(l)
	//fmt.Println(a)
}
