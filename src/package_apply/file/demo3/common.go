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
	strs := ""
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 文件末尾
			break
		}
		strs += str
	}
	writeFileName := "b.txt"
	writeFile, err := os.OpenFile(writeFileName, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)
	_, _ = writer.WriteString(strs)
	// buffer 缓存区刷新
	_ = writer.Flush()

}
