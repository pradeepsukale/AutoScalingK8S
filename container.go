package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var done chan int
var servingReq bool

func init() {
	router = gin.Default()

}

func GetuserGin(c *gin.Context) {
	c.JSON(http.StatusOK, "Hey, this is Pradeep! Have Wonderful day!!!")
}

func generateLoad(c *gin.Context) {

	if servingReq == true {
		c.JSON(http.StatusInternalServerError, "Already generating cpu load!")
		return
	}

	servingReq = true

	done = make(chan int)

	go func() {
		for {

			select {
			case <-done:
				return
			default:
			}
		}
	}()

	c.JSON(http.StatusOK, "Simulating CPU load!")

	go func() {

		time.Sleep(2 * time.Minute)
		stopLoad()

	}()

}

func stopLoad() {
	close(done)
	servingReq = false
}

func main() {
	fmt.Print("Hello World!")
	router.GET("/aboutme", GetuserGin)
	router.GET("/simulateload", generateLoad)
	router.Run(":8080")
}
