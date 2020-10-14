package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

func main() {
	cwd, err := os.Getwd()

	viper.SetConfigFile(cwd + "/src/pkg/viper_/demo1/config.yaml")
	content, err := ioutil.ReadFile(cwd + "/src/pkg/viper_/demo1/config.yaml")
	fmt.Println("err1", err)
	err = viper.ReadConfig(bytes.NewBuffer(content))
	fmt.Println("err2", err)
	a := viper.Get("app")
	fmt.Println(a)
}
