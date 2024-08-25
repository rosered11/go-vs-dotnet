package main

import (
	"demo/go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/cpu/:input", func(c *gin.Context) {
		input, _ := strconv.Atoi(c.Params.ByName("input"))
		factorial := services.NewHeavy()
		factorial.CPUBurn(input)
		c.String(http.StatusOK, "Ok")
	})
	r.GET("/memory/:input", func(c *gin.Context) {
		input, _ := strconv.Atoi(c.Params.ByName("input"))
		factorial := services.NewHeavy()
		factorial.MemoryBurn(input)
		c.String(http.StatusOK, "Ok")
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
