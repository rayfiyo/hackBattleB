package handler

import (
	"api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Title": "バンジージャンプ",
	})
}

func GetMode(c *gin.Context) {
	c.JSON(http.StatusOK, util.ReadMode)
}

func SetMode(c *gin.Context) {
	c.JSON(http.StatusOK, util.WriteMode)
}

func Clear(c *gin.Context) {
	c.JSON(http.StatusOK, util.Clear)
}
