package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"terminal/service"
	"terminal/util"
)

// Hello
// @Tags    Test
// @Summary init test function
// @Produce json
// @Router  /hello [get]
func Hello(c *gin.Context) {
	message := ""
	content, success := service.Hello()
	if success {
		message = "database initialization succeeded"
	} else {
		message = "database initialization failed"
	}
	util.UniformReturn(c, http.StatusOK, success, message, content)
}
