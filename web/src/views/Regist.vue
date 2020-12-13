<template>
<div id="Regist">

  <a class="waves-effect waves-light btn" href="/"><i class="material-icons left">arrow_back</i>首页</a>
  <div class="row login-regist-panel">
    <div class="panel-title">
      <a href="/login" class="focus-login">登录</a>
      <a>/</a>
      <a class="focus-regist">注册</a>
    </div>

    <form class="col s12">
      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">account_box</i>
          <input id="username" type="text" class="validate" v-model="username">
          <label for="username">用户名</label>
        </div>
      </div>

      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">email</i>
          <input id="email" type="email" class="validate" v-model="email">
          <label for="email">邮箱</label>
        </div>
      </div>

      <div class="row">
        <div class="input-field col s8">
          <i class="material-icons prefix">security</i>
          <input id="code" type="text" class="validate" v-model="code">
          <label for="code">验证码</label>
        </div>
        <a class="send-code-btn waves-effect waves-light btn prefix" @click="sendCode">发送验证码</a>
      </div>

      <div class="row">
        <div class="input-field col s12">
          <i class="material-icons prefix">lock</i>
          <input id="password" type="password" class="validate" v-model="password" v-on:keyup.enter="login">
          <label for="password">密码</label>
        </div>
      </div>
      <a class="waves-effect waves-light btn submit-btn" @click="regist">注册</a>
    </form>
  </div>

</div>
</template>

<script>
import '@/less/reglog.less';
import { SendCodeAPI, RegistAPI } from "@/common/api";
export default {
  name: "Regist",
  data() {
    return {
      username: '',
      email: '',
      code: '',
      password: '',
    };
  },
  methods: {
    sendCode() {
      if (this.email === '') {
        alert('邮箱不能为空!')
        return
      }
      const data = {
        email: this.email
      }
      SendCodeAPI(data).then(resp => {
        if (resp.ret !== 0){
          alert(resp.msg)
        }
      })
    },
    regist() {
      if (this.username === '') {
        alert("用户名不能为空!")
        return
      }
      if (this.email === '') {
        alert("邮箱不能为空!")
        return
      }
      if (this.code === '') {
        alert("验证码不能为空!")
        return
      }
      if (this.password === '') {
        alert("密码不能为空!")
        return
      }
      const data = {
        username: this.username,
        email: this.email,
        code: this.code,
        password: this.password
      }
      RegistAPI(data).then(resp => {
        if (resp.ret !== 0) {
          alert(resp.msg)
          return
        }
        alert("注册成功!")
      })
    }
  }
}
</script>