package model

// `文章`
type Article struct {
	ID         uint   `json:"_id" gorm:"primary_key;AUTO_INCREMENT"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime int    `json:"update_time" gorm:"autoUpdateTime"`
	CreatorID  uint   `json:"user_id" gorm:"foreignKey:UserRefer, OnDelete:CASCADE"`
}

func (Article) TableName() string {
	return "article_article"
}

// `微小说`
type Novel struct {
	ID          uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"type:string; not null"`
	Content     string `json:"content" gorm:"type:text; not null"`
	PublishTime int64  `json:"publish_time" gorm:"autoCreateTime"`
	UpdateTime  int64  `json:"update_time" gorm:"autoUpdateTime"`
	CreatorID   uint   `json:"user_id" gorm:"foreignKey:UserRefer, OnDelete:SET NULL"`
}

func (Novel) TableName() string {
	return "article_novel"
}

type Thesis struct {
	ID          uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"type:string; not null"`
	Content     string `json:"content" gorm:"type:text; not null"`
	PublishTime int64  `json:"publish_time" gorm:"autoCreateTime"`
	UpdateTime  int64  `json:"update_time" gorm:"autoUpdateTime"`
	CreatorID   uint   `json:"user_id" gorm:"foreignKey:UserRefer, OnDelete:SET NULL"`
}

func (Thesis) TableName() string {
	return "article_thesis"
}
