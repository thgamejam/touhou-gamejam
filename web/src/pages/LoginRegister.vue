<template>
  <div class="signin-background"/>
  <div class="signin-container center bg-white">
    <div class="signin-form center" v-show="isLogin">
      <div class="signin-form-account">
        <!--<p class="bold">账号</p>-->
        <input class="signin-form-account-input" placeholder="电子邮箱或手机号码">
      </div>
      <div class="signin-form-password">
        <!--<p class="bold">密码</p>-->
        <input class="signin-form-password-input" placeholder="密码">
      </div>
      <div class="signin-helper">
        <a class="signin-helper-forget-password left" href="">忘记密码</a>
        <button class="signin-helper-register right" @click="onButtonDown">立即注册</button>
      </div>

      <div class="signin-login-btn">
        <button>登入</button>
      </div>

      <el-divider style="opacity: 0.8;margin: 30px 0 30px 0;">
        <p style="font-size:5px;">或用以下方式登录</p>
      </el-divider>

      <div style="font-size: 30px; color: white">
        <i class="ri-github-fill"></i>
        <i class="ri-qq-fill" style="margin: 0 20px 0 20px;"></i>
        <i class="ri-wechat-fill"></i>
      </div>
    </div>

    <div class="signin-form center" v-show="!isLogin">
      <div class="signin-form-account">
        <!--<p class="bold">账号</p>-->
        <input class="signin-form-account-input" placeholder="电子邮箱或手机号码">
      </div>
      <div class="signin-form-password">
        <!--<p class="bold">密码</p>-->
        <input class="signin-form-password-input" placeholder="密码">
      </div>
      <div class="signin-form-password">
        <!--<p class="bold">密码</p>-->
        <input class="signin-form-password-input" placeholder="确认密码">
      </div>
      <div class="signin-helper">
        <button class="signin-helper-forget-password center">机器人验证</button>
      </div>

      <div class="signin-login-btn">
        <button>注册</button>
      </div>

      <el-divider style="opacity: 0.8;margin: 30px 0 30px 0;">
        <button style="font-size: 5px" class="signin-helper-register right" @click="onButtonDown">返回登录</button>
      </el-divider>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ref} from "vue";

let isLogin = ref(true);

function onButtonDown() {
  isLogin.value = !isLogin.value;
}
</script>

<style scoped>
input {
  padding: 0 10px;
  background-color: #fff;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
  box-sizing: border-box;
  color: #606266;
  display: inline-block;
  font-size: inherit;
  height: 100%;
  line-height: 40px;
  outline: 0;
  width: 90%;
}

i {
  color: #606266;
}

i:hover {
  color: #000000;
}

.signin-background {
  background: url('../assets/Logo_bg.png') no-repeat;
  width: 500px;
  height: 500px;
  margin: 0 auto;
}

.signin-container {
  position: absolute;
  width: 500px;
  height: 400px;
  border-radius: 20px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2);
}

.signin-form {
  transform: translate(0, 10%);
  width: 75%;
  height: 90%;
}

.signin-form-account,
.signin-form-password {
  height: 40px;
  margin: 15px auto;
}

.signin-form-account-input,
.signin-form-password-input {

}

.signin-helper {
  height: 20px;
  margin: 10px 30px;
}

.signin-helper-register,
.signin-helper-forget-password {
  cursor: pointer;
  border: transparent;
  background-color: transparent;
  font-size: 14px;
  text-decoration: none;
  color: #606266;
}

.signin-helper-register:hover,
.signin-helper-forget-password:hover {
  color: #000000;
}

.signin-login-btn {
  margin: 15px 0;
}

.signin-login-btn button {
  border: #606266 1px solid;
  border-radius: 5px;
  background-color: #606266;
  font-size: 18px;
  color: white;
  height: 40px;
  width: 90%;
}

.signin-login-btn button:hover {
  border: #000000 1px solid;
  background-color: #000000;
}

.signin-login-btn button:focus {
  border: #363636 3px solid;
  background-color: #363636;
}
</style>