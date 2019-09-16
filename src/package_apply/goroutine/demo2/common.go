package main
import (
	"fmt"
	"strconv"
)
var m = make(map[string]int)
func run(n int ) int {
	if n == 1{
		return 1
	}else{
		return n*run(n-1)
	}
		



}

func main(){
	fmt.Println(run(4))
	for i:=1; i<=200; i++{
		m[strconv.Itoa(i)] = run(i)
	}
	fmt.Println(m)
}