package main

import (
	"bufio"
	"os"
)

func main() {

	// 文件写入
	fileName := "a.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return
	}
	defer file.Close()
	str := "asdsadasd"
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(str)
	// buffer 缓存区刷新
	_ = writer.Flush()

}
