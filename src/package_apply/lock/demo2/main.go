package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
var (
	lock sync.RWMutex
	g    sync.WaitGroup
)

func read() {
	lock.RLock()
	fmt.Println("read")
	time.Sleep(1120 * time.Millisecond)
	lock.RUnlock()
	g.Done()

}
func write() {
	lock.Lock()
	fmt.Println("write")
	time.Sleep(1120 * time.Millisecond)
	lock.Unlock()
	g.Done()

}

func main() {
	for i := 1; i <= 20; i++ {
		g.Add(1)
		go read()
		g.Add(1)
		go write()
	}
	g.Wait()
}
