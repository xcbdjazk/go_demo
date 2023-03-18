package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 文件写入
	readFileName := "a.txt"
	file, err := os.Open(readFileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	strs := "123"
	for {
		str, err := reader.ReadString('\n')
		strs += str
		if err == io.EOF { // io.EOF 文件末尾
			break
		}

	}
	writeFileName := "b.txt"
	writeFile, err := os.OpenFile(writeFileName, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer writeFile.Close()
	fmt.Println(strs)
	writer := bufio.NewWriter(writeFile)
	_, _ = writer.WriteString(strs)
	// buffer 缓存区刷新
	_ = writer.Flush()

}
