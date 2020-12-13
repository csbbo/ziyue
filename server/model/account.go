package model

// `User`用户，定义了用户的基础信息
type User struct {
	ID            uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username      string `json:"username" gorm:"unique; not null"`
	Password      string `json:"password" gorm:"not null"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Remark        string `json:"remark"`
	ImgPath       string `json:"img_path"`
	CreateTime    int    `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime    int    `json:"update_time" gorm:"autoUpdateTime"`
	LastLoginTime int    `json:"last_login_time"`
}

func (User) TableName() string {
	return "account_user"
}

// `UserInfo`用户信息，与用户关联的其他重要信息
type UserInfo struct {
	Money     int   `json:"money"`
	WordCount int64 `json:"word_count"`
	LikeCount int64 `json:"like_count"`
	UserID    uint  `json:"user_id" gorm:"foreignKey:UserRefer, OnDelete:CASCADE"`
}

func (UserInfo) TableName() string {
	return "account_userinfo"
}
