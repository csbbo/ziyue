package elasticsearch

type Article struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64    `json:"create_time"`
	UpdateTime int64    `json:"update_time"`
	Uid  uint          `json:"uid"`
}

func (Article) IndexName() string {
	return "article"
}

type ArticleRes struct {
	Id string `json:"id"`
	Article
}