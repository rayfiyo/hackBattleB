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
