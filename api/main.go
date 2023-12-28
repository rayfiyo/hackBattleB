package main

import (
	"api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// r := gin.New()
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	h := handler.NewTitle()

	r.GET("/clear", handler.Clear)
	r.GET("/getMode", handler.GetMode)
	r.GET("/getTitle", h.GetTitle)
	r.GET("/setMode", handler.SetMode)
	r.GET("/sendTitle", h.SendTitle)

	r.Run()
}
