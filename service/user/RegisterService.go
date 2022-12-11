package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"terminal/define"
	"terminal/models"
	"terminal/request"
)

func ProcessRegister(c *gin.Context) error {
	// bind data
	userRegisterReq := new(request.UserRegisterReq)
	if err := c.ShouldBind(userRegisterReq); err != nil {
		return err
	}

	user := new(models.User)

	// query whether the account exists
	if err := models.DB.Where("account = ?", userRegisterReq.Account).First(user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			// system error
			return err
		}
	} else {
		// the account exists, return error information
		return errors.New("account already exists")
	}

	// if the account doesn't exist, insert it into the database
	user.Account = userRegisterReq.Account
	user.Password = userRegisterReq.Password
	user.Nickname = userRegisterReq.NickName
	user.Avatar = define.DefaultAvatar

	// insert the user information
	if err := models.DB.Create(&user).Error; err != nil {
		return err
	}

	// insert the user tag
	for i := 0; i < len(userRegisterReq.Tags); i++ {
		userTag := models.UserTag{
			UserID: user.ID,
			TagID:  userRegisterReq.Tags[i],
		}
		if err := models.DB.Create(&userTag).Error; err != nil {
			return err
		}
	}

	return nil
}
