package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"server/config"
	"server/loger"
)

func Session() gin.HandlerFunc {
	store, err := redis.NewStoreWithDB(10, "tcp", "localhost:6379", "", config.Configs.SessionDb, []byte("secret"))
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: config.Configs.SessionExpire})
	return sessions.Sessions("session", store)
}
