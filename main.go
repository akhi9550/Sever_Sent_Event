package main

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	resultChan := make(chan string)
	r.GET("/ping", func(c *gin.Context) {
		go func() {
			for i := 0; i < 15; i++ {
				resultChan <- "pong"
				time.Sleep(1 * time.Second)
			}
			close(resultChan)
		}()

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-resultChan; ok {
				c.SSEvent("Message", msg)
				return true
			}
			return false
		})
	})
	r.Run(":8080")
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.Stream(func(w io.Writer) bool {
// 			for i := 0; i < 5; i++ {
// 				c.SSEvent("message", "ping")
// 				return true
// 			}
// 			return false
// 		})
// 	})
// 	r.Run(":8080")
// }
