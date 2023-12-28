package handler

import (
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewTitle() *TitleHandler {
	return &TitleHandler{TitleInfo: &model.TitleInfo{}}
}

type TitleHandler struct {
	*model.TitleInfo
}

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

func (h *TitleHandler) GetTitle(c *gin.Context) {
	titlesArray := make([]gin.H, len(h.Titles))
	for i, t := range h.Titles {
		titlesArray[i] = gin.H{"title": t}
	}
	c.JSON(http.StatusOK, titlesArray)
}

func SetMode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}

func (h *TitleHandler) SendTitle(c *gin.Context) {
	title := c.Query("title")
	h.Titles = append(h.Titles, title)
	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})
}
