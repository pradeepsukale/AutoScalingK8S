package main

import (
	"fmt"
	"net/http"

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

	//	for i := 0; i < runtime.NumCPU(); i = i + 1 {

	go func() {
		for {

			select {
			case <-done:
				return
			default:
			}
		}
	}()
	//	}

	c.JSON(http.StatusOK, "Simulating CPU load!")
}

func stopLoad(c *gin.Context) {

	if servingReq == true {
		close(done)
		c.JSON(http.StatusOK, "Stopped CPU load!")
		servingReq = false
	} else {
		c.JSON(http.StatusOK, "Not generating any load!")
	}
}

func main() {
	fmt.Print("Hello World!")
	router.GET("/aboutme", GetuserGin)
	router.GET("/simulateload", generateLoad)
	router.GET("/stopload", stopLoad)
	router.Run(":8080")
}
