package account

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	red "github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"server/config"
	"server/db/mysql"
	"server/db/redis"
	"server/email"
	"server/loger"
	"server/model"
	"server/protocol"
	"server/utils"
	"time"
)

type Auth struct{}

func (Auth) Login(username string, password string, c *gin.Context) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	timestamp := time.Now().Unix()
	var user model.User

	db := mysql.GetConn()
	err := db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		resp.Msg = "用户不存在"
		return resp
	}

	if user.Password != utils.MkMd5(password) {
		resp.Msg = "密码错误!"
	}

	err = db.Model(&user).Where("id = ?", user.ID).Update("last_login_time", timestamp).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	user.LastLoginTime = int(timestamp)
	userStr, err := json.Marshal(user)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	session := sessions.Default(c)
	loginUser := session.Get("user")
	if loginUser != nil {
		resp.Ret = 0
		return resp
	}
	session.Set("user", string(userStr))
	err = session.Save()
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	resp.Ret = 0
	return resp
}

func (Auth) Regist(username, password, code, email string) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var user model.User

	redisConn := redis.GetConn().Get()
	cacheKey := "code_" + email
	rightCode, _ := red.String(redisConn.Do("GET", cacheKey))
	if code != rightCode {
		resp.Msg = "非常遗憾了，您的验证码不对,请再确认一下"
		return resp
	}

	db := mysql.GetConn()
	user.Username = username
	user.Password = utils.MkMd5(password)
	user.Email = email
	err := db.Create(&user).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = "哎呀！来晚了，这个用户名已经被注册了！"
		return resp
	}
	resp.Ret = 0
	return resp
}

func (Auth) SendEmail(emailAddr string) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}

	redisConn := redis.GetConn().Get()
	cacheKey := "code_" + emailAddr
	code, _ := red.String(redisConn.Do("GET", cacheKey))
	if code != "" {
		resp.Msg = "验证码已发送，请勿重复操作"
		return resp
	}
	code = utils.MkEmailCode()
	go email.Send(emailAddr, "验证码", code)
	redisConn.Do("set", cacheKey, code)
	redisConn.Do("expire", cacheKey, config.Configs.CodeExpire)
	resp.Ret = 0
	return resp
}
