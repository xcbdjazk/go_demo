package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	rep, _ := http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})
	defer rep.Body.Close()
	var p []byte
	{
		p, _ = ioutil.ReadAll(rep.Body)
		fmt.Println(string(p))
	}

	fmt.Println(string(p))

}
