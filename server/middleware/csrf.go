package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
	"server/protocol"
)

func Csrf() gin.HandlerFunc {
	resp := protocol.Resp{Ret: -1, Msg: "", Data: ""}
	return csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			resp.Msg = "CSRF token mismatch"
			c.JSON(200, resp)
			c.Abort()
		},
	})
}
