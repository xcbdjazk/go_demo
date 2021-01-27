package main

import (
	"fmt"
	"sync"
)

var g sync.WaitGroup

func a(ac <-chan int, bc chan<- int) {
	var i = 0
	for {
		if i == 100 {
			g.Done()
			return
		}
		select {
		case <-ac:
			i++
			fmt.Println("a")
			bc <- 1
		}
	}
}

func b(ac chan<- int, bc <-chan int) {
	var i = 0
	for {
		if i == 100 {
			g.Done()
			return
		}
		select {
		case <-bc:
			i++
			fmt.Println("b")
			ac <- 1
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	g.Add(1)
	go a(chA, chB)
	g.Add(1)
	go b(chA, chB)
	chA <- 1
	g.Wait()

}
