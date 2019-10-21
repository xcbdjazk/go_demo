package main

import (
	"bufio"
	_ "encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()
	//var data []byte = []byte("asd")
	//paklen := uint32(len(data))
	//var nowByte []byte
	//fmt.Println(paklen)
	//binary.BigEndian.PutUint32(nowByte[0:4], paklen)
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if strings.Trim(line, " \r\n") == "exit" {
			break
		}
		if err != nil {
			fmt.Println("err__reader", err)
			return
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("err___Write", err)
			return
		}
		for {
			var b []byte = make([]byte, 1024)
			n, err := conn.Read(b)
			if err != nil {
				fmt.Println("err__Read,", err, "退出链接")
				break
			}
			if n == 0 {
				break
			}
			fmt.Print(string(b[:n]))
		}
	}

}
