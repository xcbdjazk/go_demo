package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Open("/home/maocat/goproject/src/package_apply/flag/demo3/webwxgetvideo.mp4")
	fmt.Println(err)
	//var b []byte
	//reader := bufio.NewReader(f)
	defer func() {
		_ = f.Close()
	}()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			break
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	//f.Write(b)
	//res := bytes.NewBuffer(b)
	//fmt.Println(b)
	//fmt.Println(reader.Read(b))
	//fmt.Println(b)
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-ss", "000001.000", "-vframes", "1", "pipe:1")
	cmd.Stdin = bytes.NewBuffer(chunk)
	//stdoutBuf := &bytes.Buffer{}
	//stderrBuf := &bytes.Buffer{}
	//cmd.Stdout = stdoutBuf
	//cmd.Stderr = stderrBuf
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
	}
	fo, _ := os.Create("/home/maocat/goproject/src/package_apply/flag/demo3/1b.jpg")
	io.Copy(fo, stdout)
	err = cmd.Wait()
	if err != nil {
		fmt.Println(123)
		fmt.Println(err)
	}
	//err = cmd.Run()

	writeFileName := `/home/maocat/goproject/src/package_apply/flag/demo3/1b.jpg`
	writeFile, err := os.OpenFile(writeFileName, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)
	writer.Write(stdoutBuf.Bytes())
	writer.Flush()
	fmt.Println(reader.String())

}
