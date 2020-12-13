<template>
<div id="ArticleEdit">
  <nav-menu></nav-menu>

  <div class="article-edit">
    <div class="article-edit-left">
      <form action="#">
        <div class="file-field input-field">
          <div class="btn">
            <span>File</span>
            <input @change="uploadArticle" type="file">
          </div>
          <div class="file-path-wrapper">
            <input class="file-path validate" type="text" placeholder=".txt .docx .doc">
          </div>
        </div>
      </form>
    </div>

    <div class="article-edit-right">
      <div class="article-tool-bar">
        <button class="bar-btn" @click="publishArticle">保存</button>
        <button class="bar-btn">发布</button>
      </div>

      <div class="article-title">
        <input class="title-input" type="text" v-model="title"/>
      </div>

      <div class="article-content">
      <textarea class="content-textarea" v-model="content"></textarea>
      </div>
    </div>

    <div class="article-edit-list">
      <div class="collection" v-for="(article, index) in articleList" :key="index">
        <a href="#!" class="collection-item">{{ article.title }}</a>
<!--        <a href="#!" class="collection-item active">Alvin</a>-->
      </div>
    </div>
  </div>
</div>
</template>

<script>
import "@/less/articleedit.less"
import navMenu from "@/components/NavMenu";
import { ArticleCreateAPI, GetArticleListAPI } from "@/common/api"
import ajax from 'axios'
export default {
  name: "ArticleEdit",
  components: {
    navMenu
  },
  data() {
    return {
      title: '',
      content: '',
      articleList: [],
    };
  },
  created() {
    this.getArticleList()
  },
  methods: {
    publishArticle() {
      if (this.title === '') {
        alert("标题不能为空")
        return
      }
      if (this.title === '') {
        alert("内容不能为空")
        return
      }
      const data = {
        title: this.title,
        content: this.content,

        uploadProgress: 0,
      }
      ArticleCreateAPI(data).then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        alert("保存成功")
      })
    },
    uploadArticle(e) {
      let files = e.target.files;
      let param = new FormData();
      param.append('file', files[0])
      let headers = {
        headers:{'Content-Type':'multipart/form-data'}
      };
      ajax.post('/api/articleparse',param,headers).then(resp => {
        resp = resp.data
        for (let i=1; i<100; i++) {
          this.uploadProgress += 1
        }
        if (resp.ret === 0) {
          this.uploadProgress = 100
          this.title = resp.data.title
          this.content = resp.data.content
        } else {
          alert(files[0].name+"上传失败")
          throw "图片上传失败"
        }
      })
    },
    getArticleList() {
      GetArticleListAPI().then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        this.articleList = resp.data
      })
    }
  }
}
</script>