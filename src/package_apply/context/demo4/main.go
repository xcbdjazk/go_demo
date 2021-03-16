package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	//t :=time.After(time.Millisecond*500)
	t := time.After(time.Millisecond * 1200)
	a, _ := ctx.Value(1).(int)
	fmt.Println(a)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-t:
			fmt.Println(123)
			break LOOP
		case <-ctx.Done(): // 等待上级通知
			fmt.Println(1234)
			break LOOP
		default:
		}
	}

	wg.Done()
}

func main() {
	// 多少时间后结束
	t := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)
	ctx = context.WithValue(ctx, 1, 2)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	fmt.Println("cancel")
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
