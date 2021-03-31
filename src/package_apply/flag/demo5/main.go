package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	fi, _ := ioutil.ReadFile("dir/webwxgetvideo.mp4")
	b := runFFMPEGFromStdin(fi)
	writeImage(b)
}

func runFFMPEGFromStdin(file []byte) []byte {
	width := 360
	height := 640
	log.Print("Size of the video: ", len(file))

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vframes", "1", "-ss", "000001.000", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	cmd.Stdin = bytes.NewReader(file)

	var imageBuffer bytes.Buffer
	cmd.Stdout = &imageBuffer
	err := cmd.Run()

	if err != nil {
		log.Panic("ERROR")
	}
	return imageBuffer.Bytes()

}

func writeImage(b []byte) {
	fo, _ := os.Create("dir/1b11.jpg")
	fo.Write(b)
	fo.Close()
}
