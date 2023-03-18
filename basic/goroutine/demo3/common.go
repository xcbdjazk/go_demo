package main

import (
	"fmt"
)

func main() {
	a := make(chan int, 3)
	a <- 1
	close(a)
	<-a
	d, ok := <-a
	fmt.Println(d, ok)
	<-a

}
