package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"terminal/define"
	"time"
)

func UniformReturn(c *gin.Context, status int, success bool, message string, content interface{}) {
	c.JSON(status, gin.H{
		"success": success,
		"message": message,
		"content": content,
	})
}

func GenerateToken(id int) (string, error) {
	uc := define.UserClaim{
		ID: id,
	}
	uc.ExpiresAt = time.Now().AddDate(0, 0, 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	log.Println(token)
	tokenString, err := token.SignedString([]byte(define.JwtSecret))
	if err != nil {
		return "", err
	}
	log.Println(tokenString)
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *define.UserClaim, error) {
	uc := &define.UserClaim{}
	token, err := jwt.ParseWithClaims(tokenString, uc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(define.JwtSecret), nil
	})
	return token, uc, err
}
