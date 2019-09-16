package main

import (
	"fmt"
	"time"
)

func writeData(writeChan chan int) {
	for i := 0; i < 20; i++ {
		writeChan <- i
		fmt.Println("write --- ", i)
		time.Sleep(200 * time.Microsecond)
	}
	close(writeChan)
}
func readData(writeChan chan int, readChan chan bool) {
	for {
		i, ok := <-writeChan
		if !ok {
			break
		}
		time.Sleep(200 * time.Microsecond)
		fmt.Println("read --- ", i)

	}
	readChan <- true
	close(readChan)

}

func main() {
	writeChan := make(chan int, 20)
	readChan := make(chan bool, 1)
	go writeData(writeChan)
	go readData(writeChan, readChan)
	for {
		_, ok := <-readChan
		if !ok {
			break
		}

	}

}
