package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		d := c.Param("name")
		values := c.Request.PostForm

		fmt.Println(values)
		fmt.Println(d)
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name", func(c *gin.Context) {
		//c.FullPath() == "/user/:name/*action" // true
		//values := c.Request.Form

		fmt.Println(c.Request.Form)
		values, _ := ioutil.ReadAll(c.Request.Body)

		fmt.Println(string(values))
		c.JSON(http.StatusOK, "a")
	})

	router.Run("0.0.0.0:8080")
}
