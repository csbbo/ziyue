package protocol

import (
	"encoding/json"
	erro "errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"server/loger"
	"server/validate"
)

type Resp struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ServerErr   = "服务器错误"
	LoginRquire = "未登录"
)

func Check(request *gin.Context, data interface{}) (resp Resp, err error) {
	resp = Resp{Ret: 0, Msg: "", Data: ""}
	session := sessions.Default(request)
	user := session.Get("user")
	if user == nil {
		resp.Ret = -1
		resp.Msg = LoginRquire
		return resp, erro.New(LoginRquire)
	}

	resp, err = CheckBase(request, data)
	return resp, err
}

func CheckBase(req *gin.Context, data interface{}) (resp Resp, err error) {
	resp = Resp{Ret: 0, Msg: "", Data: ""}

	if data == nil {
		return resp, nil
	}
	if req.Request.Method == "GET" {
		err = req.ShouldBindQuery(data)
		fmt.Println(data)
	} else {
		err = req.ShouldBindJSON(data)
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(data); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		return resp, erro.New("参数验证失败")
	}

	jsonstr, _ := json.Marshal(data)
	loger.Loger.Info("Articles-Add-Params:", string(jsonstr))
	if err != nil {
		loger.Loger.Info(errors.ErrorStack(errors.Trace(err)))
		resp.Ret = -1
		resp.Msg = "参数错误"
		return resp, err
	}

	return resp, nil
}
