package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"terminal/util"
)

func SendText(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

func SendImage(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

func SendAudio(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

func Call(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

func HistoryList(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}
