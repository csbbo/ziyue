package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/handler"
	"server/protocol"
	"server/service/account"
	"server/utils"
)

type User struct{}

func (User) Get(c *gin.Context) {
	resp, err := protocol.Check(c, nil)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, err := utils.GetUser(c)
	if err != nil {
		resp.Ret = -1
		resp.Msg = protocol.ServerErr
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := account.User{}
	resp = serv.Get(&user)
	c.JSON(http.StatusOK, resp)
}

type UpdateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone" validate:"len=11"`
	Remark   string `json:"remark"`
	ImgPath  string `json:"img_path"`
}

func (User) Update(c *gin.Context) {
	var data UpdateUserParams
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	serv := account.User{}
	resp = serv.Update(data.Username, data.Password, data.Email, data.Phone, data.Remark, data.ImgPath, user.ID)
	c.JSON(http.StatusOK, resp)
}

//not use
func (User) GetList(c *gin.Context) {
	var data handler.PaginateParam
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

//not use
func (User) Delete(c *gin.Context) {
	var data handler.IDListParam
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
