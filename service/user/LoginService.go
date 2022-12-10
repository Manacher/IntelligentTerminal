package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"terminal/models"
	"terminal/request"
	"terminal/util"
)

func ProcessLogin(c *gin.Context) (string, error) {
	userLoginReq := new(request.UserLoginReq)
	if err := c.ShouldBind(&userLoginReq); err != nil {
		return "", err
	}

	// query whether both the account and password are right
	user := new(models.User)
	if err := models.DB.Where("account = ? AND password = ?", userLoginReq.Account, userLoginReq.Password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("incorrect username or password")
		} else {
			return "", err
		}
	}

	// generate token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, err
}
