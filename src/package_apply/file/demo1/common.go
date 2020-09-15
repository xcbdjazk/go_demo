package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// os 包读取
	file, err := os.Open(`file.txt`)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF { // io.EOF 文件末尾
			break
		}

	}
	//2 ioutil
	byteSlice, err1 := ioutil.ReadFile(`file.txt`)
	if err1 != nil {

	}
	fmt.Println("1233d")
	fmt.Print(string(byteSlice))

}
