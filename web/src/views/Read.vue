<template>
<div id="Read">
  <div class="read-page">
    <div class="article-read-title">{{ article.title }}</div>
    <div class="article-read-content">{{ article.content }}</div>
  </div>

</div>
</template>

<script>
import { GetArticleAPI } from "@/common/api"
import "@/less/read.less"
export default {
  name: "Read",
  data() {
    return {
      article: {},
    }
  },
  created() {
    this.getArticle()
  },
  methods: {
    getArticle() {
      const data = {
        id: this.$route.query.id
      }
      GetArticleAPI(data).then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        this.article = resp.data
      })
    }
  }
}
</script>