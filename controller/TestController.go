package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	if err := MatchInsertTest(); err != nil {
		util.UniformReturn(c, http.StatusOK, false, "test failed", err.Error())
	} else {
		util.UniformReturn(c, http.StatusOK, true, "test succeeded", "")
	}
}

// TestUpload
// @Summary used to test the file upload function
// @Param   file formData file false "the avatar image file selected by the user"
// @Router  /test/upload [post]
// @accept  multipart/form-data
func TestUpload(c *gin.Context) {
	//c.Request.FormFile("file")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}

	// calculate the file's hash mapping
	b := make([]byte, fileHeader.Size)
	fileCopy := file
	_, err = fileCopy.Read(b)
	if err != nil {
		util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
	}
	hash := fmt.Sprintf("%x", md5.Sum(b))

	// query whether the file already exists
	image := new(models.Image)
	if err := models.DB.Where("hash = ?", hash).First(image).Error; err != nil {
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				// system error
				util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
				return
			} else {
				// the file doesn't exist
				// upload the file and store its hash and url path into the database
				address, err := util.COSUpload(c)
				if err != nil {
					util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
				}

				image.Path = address
				image.Hash = hash
				if err := models.DB.Create(image).Error; err != nil {
					util.UniformReturn(c, http.StatusOK, false, err.Error(), "")
				}
				util.UniformReturn(c, http.StatusOK, true, "upload successfully", address)
			}
		}
	} else {
		// the file exists, return the path directly
		util.UniformReturn(c, http.StatusOK, true, "file exists", image.Path)
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

func MatchInsertTest() error {
	match := models.Match{
		MatcherID: 1,
		MatchedId: 1,
	}
	if err := models.DB.Create(&match).Error; err != nil {
		return err
	} else {
		return nil
	}
}
