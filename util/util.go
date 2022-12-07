package util

import (
	"github.com/gin-gonic/gin"
)

func UniformReturn(c *gin.Context, status int, success bool, message string, content interface{}) {
	c.JSON(status, gin.H{
		"success": success,
		"message": message,
		"content": content,
	})
}
