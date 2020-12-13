<template>
<div id="Login">
  <a class="waves-effect waves-light btn" href="/"><i class="material-icons left">arrow_back</i>首页</a>
  <div class="row login-regist-panel">
    <div class="panel-title">
      <a class="focus-login">登录</a>
      <a>/</a>
      <a href="/regist" class="focus-regist">注册</a>
    </div>

    <form class="col s12">
      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">account_box</i>
          <input id="username" type="text" class="validate" v-model="username" v-on:keyup.enter.native="login">
          <label for="username">用户名/邮箱</label>
        </div>
      </div>

      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">lock</i>
          <input id="password" type="password" class="validate" v-model="password" v-on:keyup.enter="login">
          <label for="password">密码</label>
        </div>
      </div>
      <a class="waves-effect waves-light btn submit-btn" @click="login">登录</a>
    </form>
  </div>

</div>
</template>

<script>
import '@/less/reglog.less';
import { LoginAPI } from "@/common/api"
export default {
  name: "Login",
  data: () => ({
    username: '',
    email: '',
    password: '',
  }),
  methods: {
    login() {
      if (this.username === '') {
        alert('用户名不能为空')
        return
      }
      if (this.password === '') {
        alert('密码不能为空')
        return
      }
      const data = {
        username: this.username,
        password: this.password
      }
      LoginAPI(data).then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        this.$router.push("/")
      })
    }
  }
}
</script>