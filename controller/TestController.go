package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"terminal/models"
	"terminal/util"
	"time"
)

// Test
// @Summary used to test function
// @Produce json
// @Router  /test [get]
func Test(c *gin.Context) {
	if err := CommentInsertTest(); err != nil {
		util.UniformReturn(c, http.StatusOK, false, "test failed", err.Error())
	} else {
		util.UniformReturn(c, http.StatusOK, true, "test succeeded", "")
	}
}

func MomentInsertTest() error {
	moment := models.Moment{
		SenderID:    1,
		TextContent: "db insert test",
		Image:       "image mock",
		ViewNum:     0,
		LikeNum:     0,
		CommentNum:  0,
	}

	if err := models.DB.Create(&moment).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AnonymousInsertTest() error {
	anonymous := models.Anonymous{
		UserId:      1,
		ExpiredTime: time.Now().Add(time.Second * 30),
	}

	if err := models.DB.Create(&anonymous).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UserTagInsertTest() error {
	userTag := models.UserTag{
		UserID: 1,
		TagID:  1,
	}
	if err := models.DB.Create(&userTag).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func CommentInsertTest() error {
	comment := models.Comment{
		MomentID:    1,
		SenderID:    1,
		ReceiverID:  1,
		BelongingID: 3,
		TextContent: "123",
		Image:       "123",
	}

	if comment.ReceiverID == 0 {
		return models.DB.Omit("ReceiverID", "BelongingID").Create(&comment).Error
	} else {
		return models.DB.Create(&comment).Error
	}

}
