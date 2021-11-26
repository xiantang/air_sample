package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// start gin

	a1, _ := ioutil.ReadFile("./a1.txt")
	b1, _ := ioutil.ReadFile("./b1.txt")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	for {
		fmt.Printf("a1: %s | b1: %s\n", string(a1), string(b1))
		time.Sleep(time.Second * 5)
	}
}
