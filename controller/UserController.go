package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"terminal/request"
	"terminal/service/user"
	"terminal/util"
)

// Login
// @Summary used to authorize user and return jwt token
// @Param   req body request.UserLoginReq true "the passed-in parameter of login function"
// @Router  /user/login [post]
func Login(c *gin.Context) {
	userLoginReq := new(request.UserLoginReq)
	if err := c.ShouldBind(&userLoginReq); err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}
	if token, err := user.Verify(userLoginReq); err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	} else {
		util.UniformReturn(c, http.StatusOK, true, "login successfully", token)
	}
}

// Register
// @Summary used to register new account
// @Param   UserRegisterReq body     request.UserRegisterReq true  "the passed-in parameter of register function"
// @Param   file            formData file                    false "the avatar image file selected by the user"
// @Router  /user/register [post]
// @accept  multipart/form-data
func Register(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}

	file, err := fileHeader.Open()
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}

	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}

	util.UniformReturn(c, http.StatusOK, true, "register successfully", b)

	//if err != nil {
	//	util.UniformReturn(c, http.StatusOK, true, "register successfully", file)
	//} else {
	//	util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	//}

	//size := file.Size
	//size += 1
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
// @Summary used to modify the user's password
// @Router  /user/password [put]
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
