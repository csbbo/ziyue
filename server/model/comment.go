package model

// `评论`
type NovelComment struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Content    string `gorm:"type:text; not null"`
	CreateTime int    `gorm:"autoCreateTime"`
	UpdateTime int    `gorm:"autoUpdateTime"`
	NovelID    uint   `gorm:"foreignKey:UserRefer, OnDelete:CASCADE"`
	CreatorID  uint   `gorm:"foreignKey:UserRefer, OnDelete:SET NULL"`
}

func (NovelComment) TableName() string {
	return "comment_novelcomment"
}
