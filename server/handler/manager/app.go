package manager

import (
	"github.com/gin-gonic/gin"
)

func SetupRoute(app *gin.Engine) {
	manager := Manager{}
	r := app.Group("/api")
	{
		r.POST("/manager", manager.PlatformMessage)
	}
}
