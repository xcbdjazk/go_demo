package main

import (
	"fmt"
	"sync"
)

var (
	max = 12000
	c1  = make(chan int, 1000)
	c2  = make(chan int, max)
	a   = sync.WaitGroup{}
)

func f1() {
	for i := 1; i <= max; i++ {
		c1 <- i
	}
	close(c1)
	a.Done()
}

func f2() {
	for value := range c1 {
		flag := true
		for c := 2; c < value; c++ {
			if value%c == 0 {
				flag = false
				break
			}
		}
		if flag {
			c2 <- value
		}
	}
	a.Done()
}

func main() {
	for i := 1; i <= 16; i++ {
		a.Add(1)
		go f2()
	}
	a.Add(1)
	f1()

	a.Wait()

	for {
		// 使用Select 不能关闭管道
		select {
		case i1 := <-c1:
			fmt.Printf("received ", i1, " from c1\n")
		case i2 := <-c2:
			fmt.Printf("sent ", i2, " to c2\n")
		default:
			fmt.Printf("no communication\n")
		}
	}
}
