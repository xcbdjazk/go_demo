package main

import "fmt"

func main() {
	a := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-a:
				fmt.Println(i)
				fmt.Println(ok)
			default:
			}
		}

	}()
	a <- 1
	a <- 2
	a <- 3
	a <- 4

}
