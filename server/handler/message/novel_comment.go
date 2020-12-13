package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db/mysql"
	"server/handler"
	"server/model"
	"server/protocol"
	"server/utils"
)

type NovelComment struct{}

type NovelCommentCreateParames struct {
	NovelID uint   `json:"novel_id" validate:"required"`
	Content string `json:"content" validate:"required"`

	CreateTime int `gorm:"autoCreateTime"`
	UpdateTime int `gorm:"autoUpdateTime"`
	CreatorID  uint
}

func (NovelComment) Create(c *gin.Context) {
	var data NovelCommentCreateParames
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	data.CreatorID = user.ID
	db := mysql.GetConn()
	err = db.Table(model.NovelComment{}.TableName()).Create(&data).Error
	if err != nil {
		resp.Ret = -1
		resp.Msg = protocol.ServerErr
	}
	c.JSON(http.StatusOK, resp)
}

type NovelCommentUpdateParames struct {
	ID      uint   `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`

	UpdateTime int `gorm:"autoUpdateTime"`
}

func (NovelComment) Update(c *gin.Context) {
	var data NovelCommentUpdateParames
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	db := mysql.GetConn()
	err = db.Table(model.NovelComment{}.TableName()).Where("id = ? and creator_id = ?", data.ID, user.ID).Updates(&data).Error
	if err != nil {
		resp.Ret = -1
		resp.Msg = protocol.ServerErr
	}
	c.JSON(http.StatusOK, resp)
}

func (NovelComment) Delete(c *gin.Context) {
	var data handler.IDOnlyParam
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	db := mysql.GetConn()
	err = db.Where("id = ? and creator_id = ?", data.ID, user.ID).Delete(model.NovelComment{}).Error
	if err != nil {
		resp.Ret = -1
		resp.Msg = protocol.ServerErr
	}
	c.JSON(http.StatusOK, resp)
}
