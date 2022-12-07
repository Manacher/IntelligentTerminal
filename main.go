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
	app.GET("/hello", controller.Hello)
	app.GET("/test", controller.Test)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
