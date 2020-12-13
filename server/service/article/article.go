package article

import (
	"context"
	"encoding/json"
	"github.com/juju/errors"
	"github.com/olivere/elastic/v7"
	"server/db/es"
	"server/elasticsearch"
	"server/loger"
	"server/protocol"
	"strconv"
	"time"
)

type Article struct{}

func (Article) Get(id string) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var article elasticsearch.Article

	els := es.GetConn()
	res, err := els.Get().Index(article.IndexName()).Id(id).Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	if !res.Found {
		resp.Msg = "文章不存在"
		return resp
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	resp.Ret = 0
	resp.Data = article
	return resp
}

func (Article) GetList(uid uint) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}

	count := 10
	offset := 0
	els := es.GetConn()
	q := elastic.NewMatchQuery("uid", uid)
	res, err := els.Search(elasticsearch.Article{}.IndexName()).
		Query(q).
		Size(count).
		From(offset).
		Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	var slice []interface{}
	for _, hit := range res.Hits.Hits {
		v := elasticsearch.ArticleRes{}
		if err := json.Unmarshal(hit.Source, &v); err == nil {
			v.Id = hit.Id
			slice = append(slice, v)
		}
	}

	resp.Ret = 0
	resp.Data = slice
	return resp
}

func (Article) Create(title, content string, userId uint) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var article elasticsearch.Article

	timestamp := time.Now().Unix()
	article.Title = title
	article.Content = content
	article.Uid = userId
	article.CreateTime = timestamp
	article.UpdateTime = timestamp


	elas := es.GetConn()
	_, err := elas.Index().
		Index(article.IndexName()).
		BodyJson(article).
		Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	resp.Ret = 0
	return resp
}

func (Article) Update(id string, title, content string, uid uint) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var article elasticsearch.Article

	query := elastic.NewBoolQuery().Must()
	queryId := elastic.NewTermQuery("_id", id)
	queryUid := elastic.NewTermQuery("uid", uid)
	query.Must(queryId, queryUid)
	timestamp := time.Now().Unix()
	str := "ctx._source['title']='" + title + "';ctx._source['content']='" + content + "';ctx._source['update_time']=" + strconv.FormatInt(timestamp, 10)
	script := elastic.NewScript(str)

	els := es.GetConn()
	_, err := els.UpdateByQuery(article.IndexName()).
		Query(query).
		Script(script).
		ProceedOnVersionConflict().
		Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	resp.Ret = 0
	return resp
}

func (Article) Delete(ids []string, uid uint) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var article elasticsearch.Article

	query := elastic.NewBoolQuery().Must()
	queryId := elastic.NewTermQuery("_id", ids) //problem
	queryUid := elastic.NewTermQuery("uid", uid)
	query.Must(queryId, queryUid)

	els := es.GetConn()
	_, err := els.DeleteByQuery(article.IndexName()).
		Query(query).
		ProceedOnVersionConflict().
		Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}
	resp.Ret = 0
	return resp
}

func (Article) ShowArticles() (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}

	count := 10
	offset := 0
	els := es.GetConn()
	res, err := els.Search(elasticsearch.Article{}.IndexName()).
		Sort("create_time", false).
		Size(count).
		From(offset).
		Do(context.Background())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		resp.Msg = protocol.ServerErr
		return resp
	}

	var slice []interface{}
	for _, hit := range res.Hits.Hits {
		v := elasticsearch.ArticleRes{}
		if err := json.Unmarshal(hit.Source, &v); err == nil {
			v.Id = hit.Id
			if len([]rune(v.Content)) > 200 {
				v.Content = string([]rune(v.Content)[:160]) + "..."
			}
			slice = append(slice, v)
		}
	}

	resp.Ret = 0
	resp.Data = slice
	return resp
}
