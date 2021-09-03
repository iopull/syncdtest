package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	fmt.Println("start ...")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("end22....")
	fmt.Println("branch buga....")

	//
	r := gin.Default()
	r.GET("/syncd/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/syncd/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"branch": "buga",
			"time":time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	r.Run("0.0.0.0:9080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
