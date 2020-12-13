package article

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
	"os"
	"path"
	"server/config"
	"server/handler"
	"server/loger"
	"server/protocol"
	"server/service/article"
	"server/utils"
)

type Article struct{}

type CreateArticleParams struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (Article) Create(c *gin.Context) {
	var data CreateArticleParams
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	serv := article.Article{}
	serv.Create(data.Title, data.Content, user.ID)
	c.JSON(http.StatusOK, resp)
}

type UpdateArticleParams struct {
	ID      string   `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (Article) Update(c *gin.Context) {
	var data UpdateArticleParams
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	serv := article.Article{}
	resp = serv.Update(data.ID, data.Title, data.Content, user.ID)
	c.JSON(http.StatusOK, resp)
}

func (Article) Get(c *gin.Context) {
	var data handler.EsIDOnlyParm
	resp, err := protocol.CheckBase(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	serv := article.Article{}
	resp = serv.Get(data.ID)
	c.JSON(http.StatusOK, resp)
}

func (Article) GetList(c *gin.Context) {
	resp, err := protocol.Check(c, nil)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	serv := article.Article{}
	resp = serv.GetList(user.ID)
	c.JSON(http.StatusOK, resp)
}

func (Article) Delete(c *gin.Context) {
	var data handler.EsIDListParm
	resp, err := protocol.Check(c, &data)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	user, _ := utils.GetUser(c)
	serv := article.Article{}
	resp = serv.Delete(data.IDS, user.ID)
	c.JSON(http.StatusOK, resp)
}

func (Article) ParseArticleFile(c *gin.Context) {
	resp, err := protocol.Check(c, nil)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	file, _ := c.FormFile("file")
	dst := utils.GetAppDir() + config.Configs.TmpDir + utils.RandString(8)
	filename := file.Filename
	ext := path.Ext(filename)
	if ext == ".txt" || ext == ".doc" || ext == ".docx" {
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
			resp.Ret = -1
			resp.Msg = protocol.ServerErr
			c.JSON(http.StatusOK, resp)
			return
		}
		title := ""
		content := ""
		if ext == ".txt" {
			title, content = utils.ReadTxt(dst)
		} else {
			title, content = utils.ReadDoc(dst)
		}
		resp.Data = map[string]string{"title": title, "content": content}
		err = os.Remove(dst)
		if err != nil {
			loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
			resp.Ret = -1
			resp.Msg = protocol.ServerErr
			c.JSON(http.StatusOK, resp)
			return
		}
	} else {
		resp.Ret = -1
		resp.Msg = "文件无法识别"
	}
	c.JSON(http.StatusOK, resp)
}

func (Article) ShowArticles(c *gin.Context) {
	serv := article.Article{}
	resp := serv.ShowArticles()
	c.JSON(http.StatusOK, resp)
}
