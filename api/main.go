package main

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
    // r := gin.New()
	r := gin.Default()

	r.GET("/getTitle", handler.GetTitle)
	r.GET("/getMode", handler.GetMode)
	r.GET("/setMode", handler.SetMode)
	r.GET("/clear", handler.Clear)

	r.Run()
}
