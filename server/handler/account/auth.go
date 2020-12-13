package account

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
	"net/http"
	"server/protocol"
	"server/service/account"
)

type Auth struct{}

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (Auth) Login(c *gin.Context) {
	var data LoginParams
	resp, err := protocol.CheckBase(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := account.Auth{}
	resp = serv.Login(data.Username, data.Password, c)
	c.JSON(http.StatusOK, resp)
}

func (Auth) Logout(c *gin.Context) {
	resp, err := protocol.Check(c, nil)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	c.JSON(http.StatusOK, resp)
}

type RegistParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (Auth) Regist(c *gin.Context) {
	var data RegistParams
	resp, err := protocol.CheckBase(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := account.Auth{}
	resp = serv.Regist(data.Username, data.Password, data.Code, data.Email)
	c.JSON(http.StatusOK, resp)
}

type SendEmailParams struct {
	Email string `json:"email" validate:"required,email"`
}

func (Auth) SendEmail(c *gin.Context) {
	var data SendEmailParams
	resp, err := protocol.CheckBase(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := account.Auth{}
	resp = serv.SendEmail(data.Email)
	c.JSON(http.StatusOK, resp)
}

func (Auth) CsrfToken(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	resp.Data = map[string]string{"X-CSRF-TOKEN": csrf.GetToken(c)}
	c.JSON(http.StatusOK, resp)
}

func (Auth) CheckAuth(c *gin.Context) {
	resp, err := protocol.Check(c, nil)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}