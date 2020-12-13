package account

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine) {
	user := User{}
	auth := Auth{}
	r := app.Group("/api")
	{
		r.GET("/user", user.Get)
		r.PUT("/user", user.Update)
		r.DELETE("/user", user.Delete)
		r.GET("/userlist", user.GetList)

		r.POST("/login", auth.Login)
		r.POST("/regist", auth.Regist)
		r.POST("/logout", auth.Logout)
		r.POST("/code", auth.SendEmail)
		r.GET("/csrftoken", auth.CsrfToken)
		r.GET("/checkauth", auth.CheckAuth)
	}
}
