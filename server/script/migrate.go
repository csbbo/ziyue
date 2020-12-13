package main

import (
	"server/config"
	"server/db/mysql"
	"server/model"
	"server/utils"
)

func init() {
	appDir := utils.GetAppDir()
	config.Default(appDir + "/conf/conf.ini")

	mysql.Init()
}

func main() {
	db := mysql.GetConn()
	db.AutoMigrate(
		&model.User{}, &model.UserInfo{},
		&model.Article{}, &model.Novel{}, &model.Thesis{},
	)
}
