package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"terminal/controller"
	"terminal/define"
	_ "terminal/docs"
)

func main() {
	app := gin.New()
	// use cors to solve the cross domain problem
	app.Use(cors.Default())
	Register(app)
	app.Run(":" + define.Port)
}

func Register(app *gin.Engine) {
	app.POST("/user/login", controller.Login)
	app.POST("/user/register", controller.Register)
	app.GET("/user/detail", controller.Detail)
	app.PUT("/user/modify", controller.Modify)
	app.PUT("/user/password", controller.Password)
	app.GET("/user/follower", controller.Follower)
	app.GET("/user/subscribed", controller.Subscribed)

	app.POST("/chat/sendText", controller.SendText)
	app.POST("/chat/sendImage", controller.SendImage)
	app.POST("/chat/sendAudio", controller.SendAudio)
	app.POST("/chat/call", controller.Call)
	app.GET("/chat/historyList", controller.HistoryList)

	app.GET("/hello", controller.Hello)

	app.GET("/test", controller.Test)
	app.POST("/test/upload", controller.TestUpload)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
