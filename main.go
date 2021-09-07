package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("start ...")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("end22....")
	fmt.Println("branch buga....")

	//
	router := gin.Default()
	router.GET("/syncd/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/syncd/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"branch": "tank 1",
			"time":time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	srv := &http.Server{
		Addr:    ":9080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
