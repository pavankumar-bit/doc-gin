package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := Setup()
	fmt.Println("Running")
	router.Run(":8088")
}

func Setup() *gin.Engine {

	router := gin.Default()
	router.GET("/", getHome)
	return router

}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
