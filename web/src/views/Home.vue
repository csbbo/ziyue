<template>
<div id="Home">
  <div class="show">
    <ul>
      <li v-for="(article, index) in articles" :key="index" @click="read(article.id)">
        <div class="title">{{ article.title }}</div>
        <div class="content">{{ article.content }}</div>
        <!--      <div class="article-info">{{ article.create_time | formatDate }}</div>-->
      </li>
    </ul>
    <button class="show-more-btn">加载更多</button>
  </div>

  <notice class="notice"></notice>
</div>
</template>

<script>
import '@/less/home.less'
import { ShowArticlesAPI } from "@/common/api"
import {formatDate} from "@/common/filter"
import Notice from "@/components/Notice"
export default {
  name: "Home",
  components: {Notice},
  data: () => ({
    articles: []
  }),
  created() {
    this.getArticles()
  },
  filters: {
    formatDate: formatDate
  },
  methods: {
    getArticles() {
      ShowArticlesAPI().then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        this.articles = resp.data
      })
    },
    read(id) {
      this.$router.push({path: '/read', query: {id: id}})
    }
  }
}
</script>