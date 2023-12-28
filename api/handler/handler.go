package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Clear(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}

func GetMode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ModeID": 0,
	})
}

func GetTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Title": "バンジージャンプ",
	})
}

func SetMode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}

func SendTitle(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		// "Status": "OK",
	})
}
