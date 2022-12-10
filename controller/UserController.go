package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"terminal/service/user"
	"terminal/util"
)

// Login
// @Summary used to authorize user and return jwt token
// @Param   req body request.UserLoginReq true "the passed-in parameter of login function"
// @Router  /user/login [post]
func Login(c *gin.Context) {
	token, err := user.ProcessLogin(c)
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
		return
	}
	util.UniformReturn(c, http.StatusOK, true, "login successfully", token)
}

// AvatarUpdate
// @Summary  used to authorize user and return jwt token
// @Param    file formData file false "the avatar image file selected by the user"
// @Router   /user/avatarUpdate [post]
// @Security ApiKeyAuth
func AvatarUpdate(c *gin.Context) {
	address, err := user.ProcessAvatarUpload(c)
	if err != nil {
		util.UniformReturn(c, http.StatusOK, true, err.Error(), "")
		return
	}
	util.UniformReturn(c, http.StatusOK, true, "upload successfully", address)
}

// Register
// @Summary used to register new account
// @Param   UserRegisterReq body request.UserRegisterReq false "the passed-in parameter of register function"
// @Router  /user/register [post]
func Register(c *gin.Context) {
	// process register pipeline
	if err := user.ProcessRegister(c); err != nil {
		util.UniformReturn(c, http.StatusOK, true, err.Error(), "")
		return
	}
	util.UniformReturn(c, http.StatusOK, true, "register successfully", "")
}

// Detail
// @Summary used to get the target user's detailed information
// @Router  /user/detail [get]
func Detail(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

// Modify
// @Summary used to modify the user's personal information
// @Router  /user/modify [put]
func Modify(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

// Password
// @Summary  used to modify the user's password
// @Router   /user/password [put]
// @Security ApiKeyAuth
func Password(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

// Follower
// @Summary used to get the user's follower list
// @Router  /user/follower [get]
func Follower(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

// Subscribed
// @Summary used to get the user's subscribed list
// @Router  /user/subscribed [get]
func Subscribed(c *gin.Context) {
	util.UniformReturn(c, http.StatusOK, true, "mock", "")
}

// TagList
// @Summary used to get all the tag
// @Router  /user/tagList [get]
func TagList(c *gin.Context) {
	tags, err := user.GetTagList()
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
		return
	}
	util.UniformReturn(c, http.StatusOK, true, "get tag list successfully", tags)
}
