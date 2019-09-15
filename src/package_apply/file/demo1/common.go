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
		if err == io.EOF { // io.EOF 文件末尾
			break
		}
		fmt.Print(str)
	}
	//2 ioutil
	byteSlice, err1 := ioutil.ReadFile(`file.txt`)
	if err1 != nil {

	}
	fmt.Print(string(byteSlice))

}
