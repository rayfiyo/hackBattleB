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

/*
func GetMode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ModeID": 0,
	})
}

func SetMode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
*/
