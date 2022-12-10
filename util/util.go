package util

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"path"
	"terminal/define"
	"terminal/models"
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

func GetUUID() string {
	return uuid.NewV4().String()
}

// COSUpload is used to upload the passed-in file into COS
func COSUpload(c *gin.Context) (string, error) {
	file, fileHeader, err := c.Request.FormFile("file")
	u, _ := url.Parse(define.BucketPath)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	key := "terminal/" + GetUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		return "", err
	}
	return define.BucketPath + "/" + key, nil
}

// VerifyFileExistence is used to verify whether the image uploaded already exists in the COS
func VerifyFileExistence(c *gin.Context) (bool, error, *models.Image) {
	image := new(models.Image)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return false, err, image
	}
	// calculate the file's hash mapping
	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		return false, err, image
	}
	hash := fmt.Sprintf("%x", md5.Sum(b))
	// query whether the file already exists
	if err := models.DB.Where("hash = ?", hash).First(image).Error; err != nil {
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				// system error
				return false, err, image
			} else {
				// the file doesn't exist
				image.Hash = hash
				return false, nil, image
			}
		}
	}
	// file exists
	return true, nil, image
}
