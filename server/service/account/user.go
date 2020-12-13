package account

import (
	"github.com/juju/errors"
	"server/db/mysql"
	"server/loger"
	"server/model"
	"server/protocol"
	"server/utils"
)

type User struct{}

type UserDetail struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Remark        string `json:"remark"`
	ImgPath       string `json:"img_path"`
	Money         int    `json:"money"`
	WordCount     int64  `json:"word_count"`
	LikeCount     int64  `json:"like_count"`
	CreateTime    int    `json:"create_time"`
	UpdateTime    int    `json:"update_time"`
	LastLoginTime int    `json:"last_login_time"`
}

func (User) Get(user *model.User) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	userInfo := model.UserInfo{}

	db := mysql.GetConn()
	err := db.Where("id = ?", user.ID).Find(user).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	err = db.Where("user_id = ?", user.ID).Find(&userInfo).Error
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	var userDetail UserDetail
	userDetail.ID = user.ID
	userDetail.Username = user.Username
	userDetail.Email = user.Email
	userDetail.Phone = user.Phone
	userDetail.Remark = user.Remark
	userDetail.ImgPath = user.ImgPath
	userDetail.Money = userInfo.Money
	userDetail.WordCount = userInfo.WordCount
	userDetail.LikeCount = userInfo.LikeCount
	userDetail.CreateTime = user.CreateTime
	userDetail.UpdateTime = user.UpdateTime
	userDetail.LastLoginTime = user.LastLoginTime

	resp.Ret = 0
	resp.Data = userDetail
	return resp
}

func (User) Update(username, password, email, phone, remark, imgPath string, updateId uint) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var updateUser model.User
	updateUser.Username = username
	updateUser.Password = utils.MkMd5(password)
	updateUser.Email = email
	updateUser.Phone = phone
	updateUser.Remark = remark
	updateUser.ImgPath = imgPath

	db := mysql.GetConn()
	err := db.Model(&updateUser).Where("id = ?", updateId).Updates(&updateUser).Error
	if err != nil {
		resp.Ret = -1
		resp.Msg = protocol.ServerErr
		return
	}

	resp.Ret = 0
	return resp
}
