package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

func yaml() {
	cwd, err := os.Getwd()

	viper.SetConfigFile("config.yaml")
	content, err := ioutil.ReadFile(cwd + "/src/pkg/viper_/demo1/config.yaml")
	fmt.Println("err1", err)
	err = viper.ReadConfig(bytes.NewBuffer(content))
	fmt.Println("err2", err)
	//a := viper.Get("app")
	v := viper.Sub("app.a")
	//s := a.(string)
	fmt.Println(v.GetInt("b"))

}

func js() {
	var m = map[string]interface{}{
		"1": "1",
		"2": "2",
		"3": map[string]string{
			"3.1": "3.2",
		},
	}
	s, _ := json.Marshal(m)

	viper.SetConfigFile("config.json")
	err := viper.ReadConfig(strings.NewReader(string(s)))
	fmt.Println("err", err)
	//a := viper.Get("app")
	v := viper.Sub("3")
	//s := a.(string)
	fmt.Println(v.GetString("3.1"))
}
func main() {
	js()
}
