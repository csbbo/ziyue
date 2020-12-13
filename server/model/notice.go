package model

type Notice struct {
	Money     int   `json:"money"`
	WordCount int64 `json:"word_count"`
	LikeCount int64 `json:"like_count"`
	UserID    uint  `json:"user_id" gorm:"foreignKey:UserRefer, OnDelete:CASCADE"`
}

func (Notice) TableName() string {
	return "notice_notice"
}

