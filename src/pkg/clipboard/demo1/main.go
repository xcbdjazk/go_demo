package main

import (
	"fmt"
	"github.com/d-tsuji/clipboard"
	"math/rand"
	"strings"
	"time"
)

var a = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func main() {
	s := g()
	fmt.Println(s)
	//复制内容到剪贴板
	fmt.Println(clipboard.Set(s))

}

func g() string {
	var builder strings.Builder
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		time.Sleep(1 * time.Nanosecond)
		c := rand.Intn(36)
		builder.WriteString(a[c])
	}
	return builder.String()
}
