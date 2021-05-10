package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func progress(ch <-chan int64) {
	for rate := range ch {
		fmt.Printf("\rrate:%3d%%", rate)
	}
}

//读取文件
func readFile() {
	rateCh := make(chan int64)
	defer close(rateCh)
	file, err := os.Open("./微信图片_20210420230926.png")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileInfo, _ := file.Stat()
	fileTotalSize := fileInfo.Size()
	go progress(rateCh)
	resultFileByte := make([]byte, 0)
	for {
		readBuf := make([]byte, 24)
		n, err := file.Read(readBuf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		resultFileByte = append(resultFileByte, readBuf...)
		time.Sleep(1 * time.Millisecond)
		go func() {
			rateCh <- int64(float32(len(resultFileByte)) / float32(fileTotalSize) * 100)
		}()
	}
	ioutil.WriteFile("./resultFile.jpg", resultFileByte, 0600)
}

func main() {
	readFile()
}
