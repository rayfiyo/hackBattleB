package main

import (
	"api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // リリースモード！！！
	r := gin.Default()

	// CORS
	r.Use(cors.Default()) // 全部のオリジン許可

	h := handler.NewTitle()
	m := handler.NewMode()

	r.GET("/clear", handler.Clear)
	r.GET("/getMode", m.GetMode)
	r.GET("/getTitle", h.GetTitle)
	r.GET("/setMode", m.SetMode)
	r.GET("/sendTitle", h.SendTitle)

	r.Run()
}
