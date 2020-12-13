<template>
<div id="Profile">
  <div class="profile-nav">
    <div class="col s12 profile-nav-center">
      <div class="profile-space"></div>

      <div class="profile-small-nav">
        <ul class="tabs profile-nav-menu">
          <li class="tab col s3"><a class="active" href="#test1">已发布</a></li>
          <li class="tab col s3"><a href="#test2">待审阅</a></li>
          <li class="tab col s3"><a href="#test4">所有</a></li>
        </ul>
      </div>
    </div>
  </div>

  <div class="profile-divider">
    <div class="profile-left">
      <img src="https://avatars1.githubusercontent.com/u/35909137?s=400&u=9dd8afe3ff9acc78ea474b5d54e879ee5a50e75c&v=4" alt="avatar" class="circle">
      <div class="profile-info">
        <div class="profile-name">{{ user.username }}</div>
        <div class="profile-remark" v-if="user.remark">{{ user.remark }}</div>
        <div class="divider"></div>
        <div class="profile-phone" v-if="user.phone"><i class="material-icons">phone</i><span>{{ user.phone }}</span></div>
        <div class="profile-email" v-if="user.email"><i class="material-icons">email</i><span>{{ user.email }}</span></div>
        <div class="profile-time">
          <span class="time-word">注册时间</span>
          <span class="time-time">{{ user.create_time | formatTime }}</span>
        </div>
        <div class="profile-time" v-if="user.last_login_time">
          <span class="time-word">最近登录</span>
          <span class="time-time">{{ user.last_login_time | formatTime }}</span>
        </div>
        <a class="btn" href="/">Edit Profile</a>
      </div>
    </div>

    <div class="profile-right">

    </div>
  </div>

</div>
</template>

<script>
import "@/less/profile.less"
import {GetUserAPI} from "@/common/api";
import {formatTime} from "@/common/filter"
export default {
  name: "Profile",
  data: () => ({
    user: ''
  }),
  mounted() {
    window.jQuery('.tabs').tabs();
  },
  created() {
    this.getUser()
  },
  methods: {
    getUser() {
      GetUserAPI().then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        this.user = resp.data
      })
    }
  },
  filters: {
    formatTime: formatTime
  },
}
</script>

<style scoped>

</style>