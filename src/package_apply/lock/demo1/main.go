package main

import (
	"fmt"
	"sync"
)

// 互斥锁
var (
	lock  sync.Mutex
	g     sync.WaitGroup
	count = 1
)

func test() {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
	g.Done()

}

func main() {
	for i := 1; i <= 20; i++ {
		g.Add(1)
		go test()
	}
	g.Wait()
}
