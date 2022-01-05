package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", HelloWorld)
	router.GET("/:userName", GetUserName)
	return router
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "World",
	})
}

func GetUserName(c *gin.Context) {

	userName := c.Param("userName")

	c.JSON(http.StatusOK, gin.H{
		"Hello": userName,
	})
}
func main() {

	var addr string
	flag.StringVar(&addr, "addr", ":8080", "web listen addr")
	flag.Parse()
	router := SetupRouter()
	router.Run(addr)
}
