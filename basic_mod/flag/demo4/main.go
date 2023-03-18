package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	fi, _ := ioutil.ReadFile("/home/maocat/goproject/src/package_apply/flag/demo4/webwxgetvideo.mp4")

	runFFMPEGFromStdin(populateStdin(fi))
}

func populateStdin(file []byte) func(io.WriteCloser) {
	fmt.Println(len(file))
	return func(stdin io.WriteCloser) {
		defer stdin.Close()
		fmt.Println(1231)
		stdin.Write(file)
		//_, err := io.Copy(stdin, bytes.NewReader(file))
		//fmt.Println(err)

	}
}

func runFFMPEGFromStdin(populate_stdin_func func(io.WriteCloser)) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-ab", "128k", "-f", "mp3", "pipe:1")
	stdin, err := cmd.StdinPipe()
	fmt.Println(123)
	if err != nil {
		log.Panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(123)
	err = cmd.Start()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(123)

	populate_stdin_func(stdin)
	fmt.Println(123)
	fo, _ := os.Create("output.mp3")
	io.Copy(fo, stdout)
	fmt.Println(123)
	err = cmd.Wait()
	if err != nil {
		log.Panic(err)
	}
}
