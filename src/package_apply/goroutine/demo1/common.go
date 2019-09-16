package main
import (
	"fmt"
	"time"
)

func run()  {
	for i:=0; i<=10;i++{
		fmt.Printf("run in %v \n", i)
		time.Sleep(time.Second)
	}
}

func main(){
	go run() //开启一个协程
	for i:=0; i<=10;i++{
		fmt.Printf("main in %v \n", i)
		time.Sleep(time.Second)
	}
}