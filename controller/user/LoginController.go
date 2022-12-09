package user

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
// @Router  /user/login [get]
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
