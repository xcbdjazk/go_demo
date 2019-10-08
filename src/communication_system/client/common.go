package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()
	var data []byte = []byte("asd")
	paklen := uint32(len(data))
	var nowByte []byte
	fmt.Println(paklen)
	binary.BigEndian.PutUint32(nowByte[0:4], paklen)
	n, err := conn.Write(nowByte)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(n)
}
