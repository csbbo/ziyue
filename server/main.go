package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"server/config"
	"server/db/es"
	"server/db/mysql"
	"server/db/redis"
	"server/email"
	"server/handler/account"
	"server/handler/article"
	"server/loger"
	"server/middleware"
	"server/utils"
)

func init() {
	appDir := utils.GetAppDir()
	config.Default(appDir + "/conf/conf.ini")

	loger.Default()
	mysql.Init()
	redis.Init()
	es.Init()
	email.Init()
}

func GetApp() *gin.Engine {
	app := gin.Default()

	app.Use(middleware.CostTime())
	app.Use(middleware.Session())
	app.Use(middleware.Csrf())

	account.SetupRoute(app)
	article.SetupRoute(app)

	return app
}

func main() {
	if config.Configs.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := GetApp()
	err := app.Run(config.Configs.HttpAccessHost + ":" + config.Configs.HttpListenPort)
	if err != nil {
		fmt.Println("start service error!!")
		os.Exit(0)
	}
}
