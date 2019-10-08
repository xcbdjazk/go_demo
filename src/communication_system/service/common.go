package main

import (
	"fmt"
	"net"
)

func doRead(con net.Conn) {
	defer con.Close()
	for {
		var b []byte
		n, err := con.Read(b)
		fmt.Println(n)
		if n != 4 || err != nil {
			fmt.Println("err,", err)
		}
		if n == 0 {
			break
		}
		fmt.Println(string(b))
	}
}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer listener.Close()
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Print(con)
		go doRead(con)
	}
}
