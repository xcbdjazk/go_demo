package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 文件copy
	readFileName := "./微信图片_20190831225715.jpg"
	file, err := os.Open(readFileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	writeFileName := `D:\goproject\src\package_apply\file\demo4\1b.jpg`
	writeFile, err := os.OpenFile(writeFileName, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)
	io.Copy(writer, reader)
	writer.Flush()

}
