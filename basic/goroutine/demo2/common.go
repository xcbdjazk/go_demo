package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var s sync.WaitGroup

func run() {
	var s1 string
	for i := 0; i < 100000; i++ {
		s1 += strconv.FormatInt(int64(i), 10)
		fmt.Println("run -", s1)
	}
	s.Done()
}

func main() {

	fmt.Println(runtime.NumCPU())
	// 设置cpu的执行单元,默认最大
	runtime.GOMAXPROCS(1)
	s.Add(1)
	go run()
	s.Add(1)
	go run()
	s.Add(1)
	go run()
	s.Wait()
	fmt.Println("main run ok")

}
