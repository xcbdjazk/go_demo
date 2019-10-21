package main

import (
	"fmt"
	"net"
)

func doRead(con net.Conn) {
	defer con.Close()
	for {
		var b []byte = make([]byte, 1024)
		n, err := con.Read(b)
		if err != nil {
			fmt.Println("err__Read,", err, "退出链接")
			break
		}
		if n == 0 {
			break
		}
		fmt.Print(string(b[:n]))
	}
	_, err := con.Write([]byte("hihihihihi"))
	if err != nil {
		fmt.Println("service error", err)
		return
	}
}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err__Listen", err)
		return
	}
	defer listener.Close()
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("err__Accept:", err)
			return
		}
		go doRead(con)
	}
}
